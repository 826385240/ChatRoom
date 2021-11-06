//***********************导入依赖的包***********************
package cmd

import (
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/handler"
	"ChatRoom/src/lib/tcpclient"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/protoout/login"
	"github.com/golang/protobuf/proto"
	"unsafe"
)

//***********************生成消息ID***********************
const (
	//下面定义 CS 后缀的消息
	MSG_ListRooms_CS      = 8193
	MSG_CreateChatRoom_CS = 8194
	MSG_JoinRoom_CS       = 8195
	MSG_LeaveRoom_CS      = 8196
	MSG_SendMessage_CS    = 8197
	MSG_SendGMString_CS   = 8198
	MSG_StartLogin_CS     = 8199
	MSG_CreateRole_CS     = 8200
	//下面定义 SC 后缀的消息
	MSG_ListRooms_SC      = 6145
	MSG_CreateChatRoom_SC = 6146
	MSG_JoinRoom_SC       = 6147
	MSG_LeaveRoom_SC      = 6148
	MSG_SendMessage_SC    = 6149
	MSG_SendGMString_SC   = 6150
	MSG_StartLogin_SC     = 6151
	MSG_CreateRole_SC     = 6152
)

//以下为各个消息段的范围
const MIN_TOFU_ID = 2048
const MAX_TOFU_ID = 4095

const MIN_CS_ID = 8192
const MAX_CS_ID = 10239

const MIN_SC_ID = 6144
const MAX_SC_ID = 8191

const MIN_TOSC_ID = 4096
const MAX_TOSC_ID = 6143

//ProtoBuffer消息占用的消息ID的上限
const MAX_PROTO_MSG_ID = 22527

//GenMsgById函数通过消息id生成对应的proto对象
func GenMsgById(msgId uint16) (proto.Message, unsafe.Pointer) {
	switch msgId {
	case MSG_ListRooms_CS:
		{
			msgPtr := &chat.MSG_ListRooms_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_ListRooms_SC:
		{
			msgPtr := &chat.MSG_ListRooms_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_CreateChatRoom_CS:
		{
			msgPtr := &chat.MSG_CreateChatRoom_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_CreateChatRoom_SC:
		{
			msgPtr := &chat.MSG_CreateChatRoom_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_JoinRoom_CS:
		{
			msgPtr := &chat.MSG_JoinRoom_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_JoinRoom_SC:
		{
			msgPtr := &chat.MSG_JoinRoom_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_LeaveRoom_CS:
		{
			msgPtr := &chat.MSG_LeaveRoom_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_LeaveRoom_SC:
		{
			msgPtr := &chat.MSG_LeaveRoom_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_SendMessage_CS:
		{
			msgPtr := &chat.MSG_SendMessage_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_SendMessage_SC:
		{
			msgPtr := &chat.MSG_SendMessage_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_SendGMString_CS:
		{
			msgPtr := &chat.MSG_SendGMString_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_SendGMString_SC:
		{
			msgPtr := &chat.MSG_SendGMString_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_StartLogin_CS:
		{
			msgPtr := &login.MSG_StartLogin_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_StartLogin_SC:
		{
			msgPtr := &login.MSG_StartLogin_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_CreateRole_CS:
		{
			msgPtr := &login.MSG_CreateRole_CS{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	case MSG_CreateRole_SC:
		{
			msgPtr := &login.MSG_CreateRole_SC{}
			return msgPtr, unsafe.Pointer(msgPtr)
		}
	}
	return nil, nil
}

//ConvertMsgById函数通过消息id生成对应的proto对象
func ConvertMsgById(msgId uint16, msgPtr unsafe.Pointer) proto.Message {
	switch msgId {
	case MSG_ListRooms_CS:
		return (*chat.MSG_ListRooms_CS)(msgPtr)
	case MSG_ListRooms_SC:
		return (*chat.MSG_ListRooms_SC)(msgPtr)
	case MSG_CreateChatRoom_CS:
		return (*chat.MSG_CreateChatRoom_CS)(msgPtr)
	case MSG_CreateChatRoom_SC:
		return (*chat.MSG_CreateChatRoom_SC)(msgPtr)
	case MSG_JoinRoom_CS:
		return (*chat.MSG_JoinRoom_CS)(msgPtr)
	case MSG_JoinRoom_SC:
		return (*chat.MSG_JoinRoom_SC)(msgPtr)
	case MSG_LeaveRoom_CS:
		return (*chat.MSG_LeaveRoom_CS)(msgPtr)
	case MSG_LeaveRoom_SC:
		return (*chat.MSG_LeaveRoom_SC)(msgPtr)
	case MSG_SendMessage_CS:
		return (*chat.MSG_SendMessage_CS)(msgPtr)
	case MSG_SendMessage_SC:
		return (*chat.MSG_SendMessage_SC)(msgPtr)
	case MSG_SendGMString_CS:
		return (*chat.MSG_SendGMString_CS)(msgPtr)
	case MSG_SendGMString_SC:
		return (*chat.MSG_SendGMString_SC)(msgPtr)
	case MSG_StartLogin_CS:
		return (*login.MSG_StartLogin_CS)(msgPtr)
	case MSG_StartLogin_SC:
		return (*login.MSG_StartLogin_SC)(msgPtr)
	case MSG_CreateRole_CS:
		return (*login.MSG_CreateRole_CS)(msgPtr)
	case MSG_CreateRole_SC:
		return (*login.MSG_CreateRole_SC)(msgPtr)
	}
	return nil
}

//***********************消息回调处理***********************
var msgCallBackHandler *handler.Handler

func InitCbHandler() *handler.Handler {
	msgCallBackHandler = handler.NewHandler()
	return msgCallBackHandler
}

func ExecCallBack(u unsafe.Pointer, m com.MsgToLogicPtr) bool {
	if m != nil {
		if m.MsgId <= MAX_PROTO_MSG_ID {
			return msgCallBackHandler.ExecHandler(m.MsgId, u, m.ProtoPtr)
		} else {
			return msgCallBackHandler.ExecHandler(m.MsgId, u, unsafe.Pointer(&m.MsgPtr.Data[0]))
		}
	}
	return false
}

const (
	EXEC_FLAG_TCPTASK   = 1
	EXEC_FLAG_TCPCLIENT = 2
)

func ConvertConnById(o com.ConnToLogicPtr) com.IBaseTcpConn {
	if o.ConnType == EXEC_FLAG_TCPTASK {
		return (*tcptask.TcpTask)(o.Conn)
	}
	if o.ConnType == EXEC_FLAG_TCPCLIENT {
		return (*tcpclient.TcpClient)(o.Conn)
	}
	return nil
}

//以下是 MSG_ListRooms_CS 相关的回调处理代码
type MSG_ListRooms_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_ListRooms_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_ListRooms_CS)
}

func (this *MSG_ListRooms_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_ListRooms_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_ListRooms_CS(f func(u *tcptask.TcpTask, m *chat.MSG_ListRooms_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_ListRooms_CS)
	var cb *MSG_ListRooms_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_ListRooms_CS)
		cb = i.(*MSG_ListRooms_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_ListRooms_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_ListRooms_CS, cb)
}

func (this *MSG_ListRooms_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_ListRooms_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_ListRooms_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_ListRooms_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_ListRooms_CS)
	var cb *MSG_ListRooms_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_ListRooms_CS)
		cb = i.(*MSG_ListRooms_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_ListRooms_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_ListRooms_CS, cb)
}

func (this *MSG_ListRooms_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_ListRooms_SC 相关的回调处理代码
type MSG_ListRooms_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_ListRooms_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_ListRooms_SC)
}

func (this *MSG_ListRooms_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_ListRooms_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_ListRooms_SC(f func(u *tcptask.TcpTask, m *chat.MSG_ListRooms_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_ListRooms_SC)
	var cb *MSG_ListRooms_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_ListRooms_SC)
		cb = i.(*MSG_ListRooms_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_ListRooms_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_ListRooms_SC, cb)
}

func (this *MSG_ListRooms_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_ListRooms_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_ListRooms_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_ListRooms_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_ListRooms_SC)
	var cb *MSG_ListRooms_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_ListRooms_SC)
		cb = i.(*MSG_ListRooms_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_ListRooms_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_ListRooms_SC, cb)
}

func (this *MSG_ListRooms_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_CreateChatRoom_CS 相关的回调处理代码
type MSG_CreateChatRoom_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_CreateChatRoom_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_CreateChatRoom_CS)
}

func (this *MSG_CreateChatRoom_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_CreateChatRoom_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_CreateChatRoom_CS(f func(u *tcptask.TcpTask, m *chat.MSG_CreateChatRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateChatRoom_CS)
	var cb *MSG_CreateChatRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateChatRoom_CS)
		cb = i.(*MSG_CreateChatRoom_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_CreateChatRoom_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_CreateChatRoom_CS, cb)
}

func (this *MSG_CreateChatRoom_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_CreateChatRoom_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_CreateChatRoom_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_CreateChatRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateChatRoom_CS)
	var cb *MSG_CreateChatRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateChatRoom_CS)
		cb = i.(*MSG_CreateChatRoom_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_CreateChatRoom_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_CreateChatRoom_CS, cb)
}

func (this *MSG_CreateChatRoom_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_CreateChatRoom_SC 相关的回调处理代码
type MSG_CreateChatRoom_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_CreateChatRoom_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_CreateChatRoom_SC)
}

func (this *MSG_CreateChatRoom_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_CreateChatRoom_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_CreateChatRoom_SC(f func(u *tcptask.TcpTask, m *chat.MSG_CreateChatRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateChatRoom_SC)
	var cb *MSG_CreateChatRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateChatRoom_SC)
		cb = i.(*MSG_CreateChatRoom_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_CreateChatRoom_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_CreateChatRoom_SC, cb)
}

func (this *MSG_CreateChatRoom_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_CreateChatRoom_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_CreateChatRoom_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_CreateChatRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateChatRoom_SC)
	var cb *MSG_CreateChatRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateChatRoom_SC)
		cb = i.(*MSG_CreateChatRoom_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_CreateChatRoom_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_CreateChatRoom_SC, cb)
}

func (this *MSG_CreateChatRoom_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_JoinRoom_CS 相关的回调处理代码
type MSG_JoinRoom_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_JoinRoom_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_JoinRoom_CS)
}

func (this *MSG_JoinRoom_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_JoinRoom_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_JoinRoom_CS(f func(u *tcptask.TcpTask, m *chat.MSG_JoinRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_JoinRoom_CS)
	var cb *MSG_JoinRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_JoinRoom_CS)
		cb = i.(*MSG_JoinRoom_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_JoinRoom_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_JoinRoom_CS, cb)
}

func (this *MSG_JoinRoom_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_JoinRoom_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_JoinRoom_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_JoinRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_JoinRoom_CS)
	var cb *MSG_JoinRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_JoinRoom_CS)
		cb = i.(*MSG_JoinRoom_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_JoinRoom_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_JoinRoom_CS, cb)
}

func (this *MSG_JoinRoom_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_JoinRoom_SC 相关的回调处理代码
type MSG_JoinRoom_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_JoinRoom_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_JoinRoom_SC)
}

func (this *MSG_JoinRoom_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_JoinRoom_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_JoinRoom_SC(f func(u *tcptask.TcpTask, m *chat.MSG_JoinRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_JoinRoom_SC)
	var cb *MSG_JoinRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_JoinRoom_SC)
		cb = i.(*MSG_JoinRoom_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_JoinRoom_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_JoinRoom_SC, cb)
}

func (this *MSG_JoinRoom_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_JoinRoom_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_JoinRoom_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_JoinRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_JoinRoom_SC)
	var cb *MSG_JoinRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_JoinRoom_SC)
		cb = i.(*MSG_JoinRoom_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_JoinRoom_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_JoinRoom_SC, cb)
}

func (this *MSG_JoinRoom_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_LeaveRoom_CS 相关的回调处理代码
type MSG_LeaveRoom_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_LeaveRoom_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_LeaveRoom_CS)
}

func (this *MSG_LeaveRoom_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_LeaveRoom_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_LeaveRoom_CS(f func(u *tcptask.TcpTask, m *chat.MSG_LeaveRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_LeaveRoom_CS)
	var cb *MSG_LeaveRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_LeaveRoom_CS)
		cb = i.(*MSG_LeaveRoom_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_LeaveRoom_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_LeaveRoom_CS, cb)
}

func (this *MSG_LeaveRoom_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_LeaveRoom_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_LeaveRoom_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_LeaveRoom_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_LeaveRoom_CS)
	var cb *MSG_LeaveRoom_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_LeaveRoom_CS)
		cb = i.(*MSG_LeaveRoom_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_LeaveRoom_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_LeaveRoom_CS, cb)
}

func (this *MSG_LeaveRoom_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_LeaveRoom_SC 相关的回调处理代码
type MSG_LeaveRoom_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_LeaveRoom_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_LeaveRoom_SC)
}

func (this *MSG_LeaveRoom_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_LeaveRoom_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_LeaveRoom_SC(f func(u *tcptask.TcpTask, m *chat.MSG_LeaveRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_LeaveRoom_SC)
	var cb *MSG_LeaveRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_LeaveRoom_SC)
		cb = i.(*MSG_LeaveRoom_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_LeaveRoom_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_LeaveRoom_SC, cb)
}

func (this *MSG_LeaveRoom_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_LeaveRoom_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_LeaveRoom_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_LeaveRoom_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_LeaveRoom_SC)
	var cb *MSG_LeaveRoom_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_LeaveRoom_SC)
		cb = i.(*MSG_LeaveRoom_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_LeaveRoom_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_LeaveRoom_SC, cb)
}

func (this *MSG_LeaveRoom_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_SendMessage_CS 相关的回调处理代码
type MSG_SendMessage_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_SendMessage_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_SendMessage_CS)
}

func (this *MSG_SendMessage_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_SendMessage_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_SendMessage_CS(f func(u *tcptask.TcpTask, m *chat.MSG_SendMessage_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendMessage_CS)
	var cb *MSG_SendMessage_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendMessage_CS)
		cb = i.(*MSG_SendMessage_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_SendMessage_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_SendMessage_CS, cb)
}

func (this *MSG_SendMessage_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_SendMessage_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_SendMessage_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_SendMessage_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendMessage_CS)
	var cb *MSG_SendMessage_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendMessage_CS)
		cb = i.(*MSG_SendMessage_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_SendMessage_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_SendMessage_CS, cb)
}

func (this *MSG_SendMessage_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_SendMessage_SC 相关的回调处理代码
type MSG_SendMessage_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_SendMessage_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_SendMessage_SC)
}

func (this *MSG_SendMessage_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_SendMessage_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_SendMessage_SC(f func(u *tcptask.TcpTask, m *chat.MSG_SendMessage_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendMessage_SC)
	var cb *MSG_SendMessage_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendMessage_SC)
		cb = i.(*MSG_SendMessage_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_SendMessage_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_SendMessage_SC, cb)
}

func (this *MSG_SendMessage_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_SendMessage_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_SendMessage_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_SendMessage_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendMessage_SC)
	var cb *MSG_SendMessage_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendMessage_SC)
		cb = i.(*MSG_SendMessage_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_SendMessage_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_SendMessage_SC, cb)
}

func (this *MSG_SendMessage_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_SendGMString_CS 相关的回调处理代码
type MSG_SendGMString_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_SendGMString_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_SendGMString_CS)
}

func (this *MSG_SendGMString_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_SendGMString_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_SendGMString_CS(f func(u *tcptask.TcpTask, m *chat.MSG_SendGMString_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendGMString_CS)
	var cb *MSG_SendGMString_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendGMString_CS)
		cb = i.(*MSG_SendGMString_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_SendGMString_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_SendGMString_CS, cb)
}

func (this *MSG_SendGMString_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_SendGMString_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_SendGMString_CS(f func(u *tcpclient.TcpClient, m *chat.MSG_SendGMString_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendGMString_CS)
	var cb *MSG_SendGMString_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendGMString_CS)
		cb = i.(*MSG_SendGMString_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_SendGMString_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_SendGMString_CS, cb)
}

func (this *MSG_SendGMString_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_SendGMString_SC 相关的回调处理代码
type MSG_SendGMString_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *chat.MSG_SendGMString_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *chat.MSG_SendGMString_SC)
}

func (this *MSG_SendGMString_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*chat.MSG_SendGMString_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_SendGMString_SC(f func(u *tcptask.TcpTask, m *chat.MSG_SendGMString_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendGMString_SC)
	var cb *MSG_SendGMString_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendGMString_SC)
		cb = i.(*MSG_SendGMString_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_SendGMString_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_SendGMString_SC, cb)
}

func (this *MSG_SendGMString_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*chat.MSG_SendGMString_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_SendGMString_SC(f func(u *tcpclient.TcpClient, m *chat.MSG_SendGMString_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_SendGMString_SC)
	var cb *MSG_SendGMString_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_SendGMString_SC)
		cb = i.(*MSG_SendGMString_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_SendGMString_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_SendGMString_SC, cb)
}

func (this *MSG_SendGMString_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_StartLogin_CS 相关的回调处理代码
type MSG_StartLogin_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *login.MSG_StartLogin_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *login.MSG_StartLogin_CS)
}

func (this *MSG_StartLogin_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*login.MSG_StartLogin_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_StartLogin_CS(f func(u *tcptask.TcpTask, m *login.MSG_StartLogin_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_StartLogin_CS)
	var cb *MSG_StartLogin_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_StartLogin_CS)
		cb = i.(*MSG_StartLogin_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_StartLogin_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_StartLogin_CS, cb)
}

func (this *MSG_StartLogin_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*login.MSG_StartLogin_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_StartLogin_CS(f func(u *tcpclient.TcpClient, m *login.MSG_StartLogin_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_StartLogin_CS)
	var cb *MSG_StartLogin_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_StartLogin_CS)
		cb = i.(*MSG_StartLogin_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_StartLogin_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_StartLogin_CS, cb)
}

func (this *MSG_StartLogin_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_StartLogin_SC 相关的回调处理代码
type MSG_StartLogin_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *login.MSG_StartLogin_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *login.MSG_StartLogin_SC)
}

func (this *MSG_StartLogin_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*login.MSG_StartLogin_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_StartLogin_SC(f func(u *tcptask.TcpTask, m *login.MSG_StartLogin_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_StartLogin_SC)
	var cb *MSG_StartLogin_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_StartLogin_SC)
		cb = i.(*MSG_StartLogin_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_StartLogin_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_StartLogin_SC, cb)
}

func (this *MSG_StartLogin_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*login.MSG_StartLogin_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_StartLogin_SC(f func(u *tcpclient.TcpClient, m *login.MSG_StartLogin_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_StartLogin_SC)
	var cb *MSG_StartLogin_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_StartLogin_SC)
		cb = i.(*MSG_StartLogin_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_StartLogin_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_StartLogin_SC, cb)
}

func (this *MSG_StartLogin_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_CreateRole_CS 相关的回调处理代码
type MSG_CreateRole_CS_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *login.MSG_CreateRole_CS)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *login.MSG_CreateRole_CS)
}

func (this *MSG_CreateRole_CS_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*login.MSG_CreateRole_CS)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_CreateRole_CS(f func(u *tcptask.TcpTask, m *login.MSG_CreateRole_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateRole_CS)
	var cb *MSG_CreateRole_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateRole_CS)
		cb = i.(*MSG_CreateRole_CS_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_CreateRole_CS_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_CreateRole_CS, cb)
}

func (this *MSG_CreateRole_CS_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*login.MSG_CreateRole_CS)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_CreateRole_CS(f func(u *tcpclient.TcpClient, m *login.MSG_CreateRole_CS)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateRole_CS)
	var cb *MSG_CreateRole_CS_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateRole_CS)
		cb = i.(*MSG_CreateRole_CS_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_CreateRole_CS_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_CreateRole_CS, cb)
}

func (this *MSG_CreateRole_CS_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}

//以下是 MSG_CreateRole_SC 相关的回调处理代码
type MSG_CreateRole_SC_CB struct {
	execFlag          int
	TcpTaskRealExec   func(u *tcptask.TcpTask, m *login.MSG_CreateRole_SC)
	TcpClientRealExec func(u *tcpclient.TcpClient, m *login.MSG_CreateRole_SC)
}

func (this *MSG_CreateRole_SC_CB) TcpTaskExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcptask.TcpTask)(u)
	msgPtr := (*login.MSG_CreateRole_SC)(m)
	this.TcpTaskRealExec(exePtr, msgPtr)
}

func Reg_TcpTask_MSG_CreateRole_SC(f func(u *tcptask.TcpTask, m *login.MSG_CreateRole_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateRole_SC)
	var cb *MSG_CreateRole_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateRole_SC)
		cb = i.(*MSG_CreateRole_SC_CB)
		cb.TcpTaskRealExec = f
	} else {
		cb = &MSG_CreateRole_SC_CB{TcpTaskRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPTASK
	msgCallBackHandler.RegHandler(MSG_CreateRole_SC, cb)
}

func (this *MSG_CreateRole_SC_CB) TcpClientExec(u unsafe.Pointer, m unsafe.Pointer) {
	exePtr := (*tcpclient.TcpClient)(u)
	msgPtr := (*login.MSG_CreateRole_SC)(m)
	this.TcpClientRealExec(exePtr, msgPtr)
}

func Reg_TcpClient_MSG_CreateRole_SC(f func(u *tcpclient.TcpClient, m *login.MSG_CreateRole_SC)) {
	bExist := msgCallBackHandler.IsExist(MSG_CreateRole_SC)
	var cb *MSG_CreateRole_SC_CB = nil
	if bExist {
		i := msgCallBackHandler.GetHandler(MSG_CreateRole_SC)
		cb = i.(*MSG_CreateRole_SC_CB)
		cb.TcpClientRealExec = f
	} else {
		cb = &MSG_CreateRole_SC_CB{TcpClientRealExec: f}
	}
	cb.execFlag = cb.execFlag | EXEC_FLAG_TCPCLIENT
	msgCallBackHandler.RegHandler(MSG_CreateRole_SC, cb)
}

func (this *MSG_CreateRole_SC_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {
	if this.execFlag&EXEC_FLAG_TCPTASK > 0 {
		this.TcpTaskExec(u, m)
		return
	}
	if this.execFlag&EXEC_FLAG_TCPCLIENT > 0 {
		this.TcpClientExec(u, m)
		return
	}
}
