package server

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/exception"
	"ChatRoom/src/lib/handler"
	"ChatRoom/src/lib/ipaddr"
	"ChatRoom/src/lib/logger"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/message"
	"ChatRoom/src/lib/netserver"
	"ChatRoom/src/lib/serialize"
	"ChatRoom/src/lib/tcpclient"
	"ChatRoom/src/lib/tcptask"
	"net"
	"unsafe"
)

const MAX_NEW_TASK_CHAN_ELEMENTS = 100000

type Server struct {
	logicPtr  *logic.Logic
	inListen  net.Listener
	outListen net.Listener
	*netserver.BaseNetServer
	cmdHandles []handler.ICmdHandle
	chMsgs     chan com.MsgToLogicPtr
}

func NewServer(innerProto string, innerIp string, innerPort uint16, outerProto string, outerIp string, outerPort uint16) *Server {
	base := netserver.NewBaseNetServer(innerProto, innerIp, innerPort, outerProto, outerIp, outerPort)
	return &Server{logicPtr: logic.InitLogic(nil), BaseNetServer: base, chMsgs: make(chan com.MsgToLogicPtr, com.MAX_TASK_DATA_CHAN_ELEMENTS)}
}

func (this *Server) AddCmdHandle(h handler.ICmdHandle) {
	this.cmdHandles = append(this.cmdHandles, h)
}

func (this *Server) RegCmdHandles() {
	for _, h := range this.cmdHandles {
		h.Init()
	}
}

func (this *Server) InnerListen() (net.Listener, error) {
	var err error
	this.inListen, err = this.BaseNetServer.InnerListen()
	if err != nil {
		exception.ErrorExit(err)
	}
	return this.inListen, err
}

func (this *Server) OuterListen() (net.Listener, error) {
	var err error
	this.outListen, err = this.BaseNetServer.OuterListen()
	if err != nil {
		exception.ErrorExit(err)
	}
	return this.outListen, err
}

func (this *Server) Connect(netProto string, ip string, port uint16) *tcpclient.TcpClient {
	addr := ipaddr.IpAddr{netProto, ip, port}
	conn, err := addr.Connect()
	if err != nil {
		logger.PANIC("错误!连接对端错误[协议:%s,IP:%s,端口:%d],%s", netProto, ip, port, err.Error())
		return nil
	}
	newclient := tcpclient.NewTcpClient(conn)
	if newclient != nil {
		newclient.WgAdd(1)
		newclient.SetWChan(this.chMsgs)
		this.logicPtr.AddNewConn(&com.ConnToLogic{com.CONN_TYPE_CLIENT, unsafe.Pointer(newclient)})

		go this.GetMsgFromConn(newclient)
		go this.PostMsgToConn(newclient)
		return newclient
	}
	return nil
}

func (this *Server) GetMsgFromConn(conn com.IBaseTcpConn) {
	conn.WgWait()
	for {
		var msg *message.Message = conn.RecvCmd()
		if msg == nil {
			//通知主逻辑协程连接关闭
			conn.SendToWChan(&com.MsgToLogic{conn.GetUniqId(), 0, nil, nil})
			return
		}

		//处理消息
		if msg.MsgId <= cmd.MAX_PROTO_MSG_ID {
			//protobuffer消息反序列化
			protoMsg, protoPtr := cmd.GenMsgById(msg.MsgId)
			if protoMsg != nil {
				serialize.UnserializeToBuffer(msg.Data, protoMsg)
				conn.SendToWChan(&com.MsgToLogic{conn.GetUniqId(), msg.MsgId, msg, protoPtr})
			}
		} else {
			//其他消息
			conn.SendToWChan(&com.MsgToLogic{conn.GetUniqId(), msg.MsgId, msg, nil})
		}
	}
}

func (this *Server) PostMsgToConn(conn com.IBaseTcpConn) {
	conn.WgWait()
	defer conn.CloseConn()
	defer conn.CloseWChan()
	defer conn.CloseRChan()
	for {
		ct := conn.RecvFromRChan()
		if ct != nil {
			if ct.MsgId <= cmd.MAX_PROTO_MSG_ID {
				//protobuffer消息序列化
				if ct.ProtoPtr != nil {
					data, _ := serialize.SerializeToBuffer(cmd.ConvertMsgById(ct.MsgId, ct.ProtoPtr))
					if !conn.SendCmd(ct.MsgId, data) {
						break
					}
				}
			} else {
				//其他消息
				if ct.MsgPtr != nil && !conn.SendCmd(ct.MsgPtr.MsgId, ct.MsgPtr.Data) {
					break
				}
			}
		} else {
			break
		}
	}
}

func (this *Server) listenTask(liconn net.Listener, ch chan *tcptask.TcpTask) {
	if liconn == nil {
		return
	}

	for {
		//接收新的连接
		conn, err := this.Accept(liconn)
		if err != nil {
			logger.DEBUG("错误!接受连接错误!")
			return
		}

		task := tcptask.NewTcpTask(conn)
		if task == nil {
			logger.DEBUG("错误!创建TcpTask错误!")
			return
		}

		//将每个task发送主线程
		task.WgAdd(1)
		if task != nil {
			ch <- task
		}
	}
}

func (this *Server) Start(Ontimer func(time int64)) bool {
	if Ontimer == nil || this.logicPtr == nil {
		return false
	}
	//初始化消息回调模块
	cmd.InitCbHandler()
	this.RegCmdHandles()

	//启动逻辑协程
	chConns := make(chan com.ConnToLogicPtr, com.MAX_LOGIC_CONNS_CHAN_ELEMENTS)
	this.logicPtr.SetOnTimer(Ontimer)
	go this.logicPtr.ServerLogic(chConns, this.chMsgs)

	//启动监听连接协程
	this.InnerListen()
	this.OuterListen()
	chNewTask := make(chan *tcptask.TcpTask, MAX_NEW_TASK_CHAN_ELEMENTS)
	go this.listenTask(this.inListen, chNewTask)
	go this.listenTask(this.outListen, chNewTask)

	for {
		newTask := <-chNewTask
		if newTask != nil {
			newTask.SetWChan(this.chMsgs)
			go this.GetMsgFromConn(newTask)
			go this.PostMsgToConn(newTask)

			//将新的task传送给逻辑协程
			chConns <- &com.ConnToLogic{com.CONN_TYPE_TASK, unsafe.Pointer(newTask)}
		}
	}
	return true
}
