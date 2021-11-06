package logic

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/logger"
	"time"
	"unsafe"
)

const LOGIC_SLEEP_TIME = int64(10 * time.Millisecond)

var g_logic *Logic

func SendMsg(c com.IBaseTcpConn, msgId uint16, msg unsafe.Pointer) {
	o := &com.MsgToLogic{c.GetUniqId(), msgId, nil, msg}
	if !c.SendToRChanNB(o) {
		g_logic.AppendMsgToQueue(o)
	}
}

type Logic struct {
	onTimer   func(tm int64)
	connMap   map[uint64]com.ConnToLogicPtr
	sendQueue []com.MsgToLogicPtr
}

func NewLogic(OnTimer func(time int64)) *Logic {
	return &Logic{onTimer: OnTimer, connMap: make(map[uint64]com.ConnToLogicPtr)}
}

func InitLogic(OnTimer func(time int64)) *Logic {
	g_logic = NewLogic(OnTimer)
	if g_logic == nil {
		logger.PANIC("错误!默认logic创建错误!")
		return nil
	}
	return g_logic
}

func (this *Logic) SetOnTimer(onTimer func(tm int64)) {
	this.onTimer = onTimer
}

func (this *Logic) AppendMsgToQueue(o com.MsgToLogicPtr) {
	this.sendQueue = append(this.sendQueue, o)
}

func (this *Logic) PostAllMsg() bool {
	length := len(this.sendQueue)
	if length <= 0 {
		return true
	}

	//还有消息堆积没有处理
	for i := 0; i < length; i++ {
		o := this.sendQueue[i]
		var c com.IBaseTcpConn = cmd.ConvertConnById(this.connMap[o.UniqId])
		if c == nil || c.SendToRChanNB(o) {
			//连接不存在或者发送成功,清除消息
			this.sendQueue[i] = nil
		}
	}

	//整理没处理完的消息
	leftMsgs := []com.MsgToLogicPtr{}
	for i := 0; i < length; i++ {
		if this.sendQueue[i] != nil {
			leftMsgs = append(leftMsgs, this.sendQueue[i])
		}
	}
	this.sendQueue = leftMsgs
	return true
}

func (this *Logic) AddNewConn(conn com.ConnToLogicPtr) {
	if conn.Conn != nil {
		uniqId := *(*uint64)(conn.Conn)
		this.connMap[uniqId] = conn

		cmd.ConvertConnById(conn).WgDone()
	}
}

func (this *Logic) DispatchMsg(t com.MsgToLogicPtr) bool {
	if conn, ok := this.connMap[t.UniqId]; ok {
		if t.MsgPtr == nil && t.ProtoPtr == nil {
			//连接关闭处理
			var c com.IBaseTcpConn = cmd.ConvertConnById(conn)
			if c != nil {
				c.SendToRChanNB(nil)
			}
			delete(this.connMap, t.UniqId)
		} else {
			cmd.ExecCallBack(conn.Conn, t)
		}
		return true
	}
	return false
}

func (this *Logic) ServerLogic(chLogic chan com.ConnToLogicPtr, chMsgs chan com.MsgToLogicPtr) {
	for {
		select {
		case c := <-chLogic:
			//从主线程获得新的连接消息
			this.AddNewConn(c)
		case t := <-chMsgs:
			//循环处理连接消息
			this.DispatchMsg(t)
		default:
			curTime := time.Now().Unix()
			this.onTimer(curTime)
			this.PostAllMsg()

			//逻辑协程睡眠
			now := time.Now().Unix()
			if now < curTime+LOGIC_SLEEP_TIME {
				time.Sleep(time.Duration(curTime + LOGIC_SLEEP_TIME - now))
			}
		}
	}
}

func (this *Logic) ClientLogic(chMsgs chan com.MsgToLogicPtr) {
	for {
		select {
		case t := <-chMsgs:
			//循环处理连接消息
			this.DispatchMsg(t)
		default:
			curTime := time.Now().Unix()
			this.onTimer(curTime)
			this.PostAllMsg()

			//逻辑协程睡眠
			now := time.Now().Unix()
			if now < curTime+LOGIC_SLEEP_TIME {
				time.Sleep(time.Duration(curTime + LOGIC_SLEEP_TIME - now))
			}
		}
	}
}
