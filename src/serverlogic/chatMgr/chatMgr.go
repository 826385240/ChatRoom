package chatMgr

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/protoout/login"
	"ChatRoom/src/serverlogic/roleMgr"
	"ChatRoom/src/serverlogic/roomMgr"
	"unsafe"
)

type ChatManager struct {
}

func (this *ChatManager) Init() {
	cmd.Reg_TcpTask_MSG_ListRooms_CS(this.MSG_ListRooms_CS_CB)
	cmd.Reg_TcpTask_MSG_CreateChatRoom_CS(this.MSG_CreateChatRoom_CS_CB)
	cmd.Reg_TcpTask_MSG_JoinRoom_CS(this.MSG_JoinRoom_CS_CB)
	cmd.Reg_TcpTask_MSG_LeaveRoom_CS(this.MSG_LeaveRoom_CS_CB)
	cmd.Reg_TcpTask_MSG_SendMessage_CS(this.MSG_SendMessage_CS_CB)
	cmd.Reg_TcpTask_MSG_SendGMString_CS(this.MSG_SendGMString_CS_CB)
}

func (this *ChatManager) MSG_ListRooms_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_ListRooms_CS) {
	ret := &chat.MSG_ListRooms_SC{}
	ret.Rooms = roomMgr.RoomManager.GetAllRooms()
	logic.SendMsg(task, cmd.MSG_ListRooms_SC, unsafe.Pointer(ret))
}

func (this *ChatManager) MSG_CreateChatRoom_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_CreateChatRoom_CS) {
	ret := &chat.MSG_CreateChatRoom_SC{}
	ret.Retcode = roomMgr.RoomManager.CreateRoom(msg.Name)
	logic.SendMsg(task, cmd.MSG_CreateChatRoom_SC, unsafe.Pointer(ret))
}

func (this *ChatManager) MSG_JoinRoom_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_JoinRoom_CS) {
	role := roleMgr.RoleManager.GetRoleName(task)
	if role == "" {
		ret := &chat.MSG_JoinRoom_SC{Retcode: false}
		logic.SendMsg(task, cmd.MSG_JoinRoom_SC, unsafe.Pointer(ret))
		return
	}

	if roomMgr.RoomManager.JoinRoom(task, role, msg.Name) == false {
		ret := &chat.MSG_JoinRoom_SC{Retcode: false}
		logic.SendMsg(task, cmd.MSG_JoinRoom_SC, unsafe.Pointer(ret))
		return
	}
}

func (this *ChatManager) MSG_LeaveRoom_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_LeaveRoom_CS) {
	role := roleMgr.RoleManager.GetRoleName(task)
	if role == "" {
		ret := &chat.MSG_LeaveRoom_SC{Retcode: false}
		logic.SendMsg(task, cmd.MSG_LeaveRoom_SC, unsafe.Pointer(ret))
		return
	}
	retcode := roomMgr.RoomManager.LeaveRoom(role)

	ret := &chat.MSG_LeaveRoom_SC{Retcode: retcode}
	logic.SendMsg(task, cmd.MSG_LeaveRoom_SC, unsafe.Pointer(ret))
}

func (this *ChatManager) MSG_SendMessage_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_SendMessage_CS) {
	role := roleMgr.RoleManager.GetRoleName(task)
	if role == "" {
		ret := &chat.MSG_SendMessage_CS{}
		logic.SendMsg(task, cmd.MSG_SendMessage_CS, unsafe.Pointer(ret))
		return
	}

	roomMgr.RoomManager.SendMessage(role, msg.Message)
}

func (this *ChatManager) MSG_SendGMString_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_SendGMString_CS) {
	ret := &login.MSG_StartLogin_SC{}
	logic.SendMsg(task, cmd.MSG_StartLogin_SC, unsafe.Pointer(ret))
}
