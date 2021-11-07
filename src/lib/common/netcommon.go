package com

import (
	"ChatRoom/src/lib/message"
	"unsafe"
)

//主线程与逻辑协程间传递连接的channel缓冲大小
const MAX_LOGIC_CONNS_CHAN_ELEMENTS = 5000

//Task网络协程向逻辑协程间传递数据的channel缓冲大小
const MAX_TASK_DATA_CHAN_ELEMENTS = 65535
const MAX_CLIENT_DATA_CHAN_ELEMENTS = 65535

//连接的读channel的缓冲区大小
const READ_CHAN_BUFFER_SIZE = 2048

//连接的写channel的缓冲区大小
const WRITE_CHAN_BUFFER_SIZE = 2048

type ConnToLogic struct {
	ConnType uint16
	Conn     unsafe.Pointer
}
type ConnToLogicPtr *ConnToLogic

type MsgToLogic struct {
	UniqId   uint64
	MsgId    uint16
	MsgPtr   *message.Message
	ProtoPtr unsafe.Pointer
}

type MsgToLogicPtr *MsgToLogic

type IBaseTcpConn interface {
	WgDone()
	WgWait()
	WgAdd(i int)
	GetUniqId() uint64
	RecvCmd() *message.Message
	SendCmd(msgId uint16, data []byte) bool
	SendToRChan(d MsgToLogicPtr)
	SendToWChan(d MsgToLogicPtr)
	SendToRChanNB(d MsgToLogicPtr) bool
	SendToWChanNB(d MsgToLogicPtr) bool
	RecvFromRChan() MsgToLogicPtr
	RecvFromWChan() MsgToLogicPtr
	RecvFromRChanNB() MsgToLogicPtr
	RecvFromWChanNB() MsgToLogicPtr
	CloseConn()
	CloseWChan()
	CloseRChan()
	IsWChanValid() bool
	IsRChanValid() bool
}
