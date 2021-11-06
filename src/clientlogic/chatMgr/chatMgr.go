package chatMgr

import (
	"ChatRoom/src/clientlogic/operateMgr"
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcpclient"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/protoout/login"
	"fmt"
	"unsafe"
)

type ChatManager struct {
}

func (this *ChatManager) Init() {
	cmd.Reg_TcpClient_MSG_ListRooms_SC(this.MSG_ListRooms_SC_CB)
	cmd.Reg_TcpClient_MSG_CreateChatRoom_SC(this.MSG_CreateChatRoom_SC_CB)
	cmd.Reg_TcpClient_MSG_JoinRoom_SC(this.MSG_JoinRoom_SC_CB)
	cmd.Reg_TcpClient_MSG_LeaveRoom_SC(this.MSG_LeaveRoom_SC_CB)
	cmd.Reg_TcpClient_MSG_SendMessage_SC(this.MSG_SendMessage_SC_CB)
	cmd.Reg_TcpClient_MSG_SendGMString_SC(this.MSG_SendGMString_SC_CB)
}

func (this *ChatManager) MSG_ListRooms_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_ListRooms_SC) {
	if len(msg.Rooms)<=0 {
		fmt.Println("当前还没有聊天室!")
	} else {
		fmt.Println("以下是正在开放聊天室:")
		for i:=0; i<len(msg.Rooms); i++ {
			fmt.Println(msg.Rooms[i])
		}
	}
	operateMgr.OperateMgr.ShowOperate(client)
}

func (this *ChatManager) MSG_CreateChatRoom_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_CreateChatRoom_SC) {
	if msg.Retcode {
		fmt.Println("创建聊天室成功!")
	} else {
		fmt.Println("创建聊天室失败,聊天室可能已经存在!")
	}
	operateMgr.OperateMgr.ShowOperate(client)
}

func (this *ChatManager) MSG_JoinRoom_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_JoinRoom_SC) {
	if msg.Retcode {
		fmt.Println("加入聊天室成功!")
	} else {
		fmt.Println("加入聊天室失败,聊天室可能不存在!")
	}
	operateMgr.OperateMgr.ShowOperate(client)
}

func (this *ChatManager) MSG_LeaveRoom_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_LeaveRoom_SC) {
	if msg.Retcode {
		fmt.Println("离开聊天室成功!")
	} else {
		fmt.Println("离开聊天室失败,聊天室可能不存在!")
	}
	operateMgr.OperateMgr.ShowOperate(client)
}

func (this *ChatManager) MSG_SendMessage_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_SendMessage_SC) {
	for i:=0; i<len(msg.Message); i++ {
		fmt.Println(msg.Message[i])
	}

	operateMgr.OperateMgr.ShowOperate(client)
}

func (this *ChatManager) MSG_SendGMString_SC_CB(client *tcpclient.TcpClient, msg *chat.MSG_SendGMString_SC) {
	ret := &login.MSG_CreateRole_CS{}
	logic.SendMsg(client, cmd.MSG_CreateRole_CS, unsafe.Pointer(ret))
}
