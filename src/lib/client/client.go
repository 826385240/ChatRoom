package client

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/handler"
	"ChatRoom/src/lib/logger"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/message"
	"ChatRoom/src/lib/netclient"
	"ChatRoom/src/lib/serialize"
	"ChatRoom/src/lib/tcpclient"
	"unsafe"
)

type Client struct {
	logicPtr *logic.Logic
	netclient.BaseNetClient
	cmdHandles []handler.ICmdHandle
	chMsgs     chan com.MsgToLogicPtr
}

func NewClient() *Client {
	return &Client{logicPtr: logic.InitLogic(nil), chMsgs: make(chan com.MsgToLogicPtr, com.MAX_CLIENT_DATA_CHAN_ELEMENTS)}
}

func (this *Client) AddCmdHandle(h handler.ICmdHandle) {
	this.cmdHandles = append(this.cmdHandles, h)
}

func (this *Client) RegCmdHandles() {
	for _, h := range this.cmdHandles {
		h.Init()
	}
}

func (this *Client) Connect(netProto string, ip string, port uint16) *tcpclient.TcpClient {
	conn, err := this.BaseNetClient.Connect(netProto, ip, port)
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

func (this *Client) GetMsgFromConn(conn com.IBaseTcpConn) {
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

func (this *Client) PostMsgToConn(conn com.IBaseTcpConn) {
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

func (this *Client) Start(OnTimer func(tm int64)) bool {
	if OnTimer == nil || this.logicPtr == nil {
		return false
	}

	//初始化消息回调模块
	cmd.InitCbHandler()
	this.RegCmdHandles()

	//启动逻辑协程
	this.logicPtr.SetOnTimer(OnTimer)
	this.logicPtr.ClientLogic(this.chMsgs)
	return true
}
