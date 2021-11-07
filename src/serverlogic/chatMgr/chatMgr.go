package chatMgr

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/serverlogic/roleMgr"
	"ChatRoom/src/serverlogic/roomMgr"
	"strconv"
	"strings"
	"time"
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

	//获取所有的房间名
	ret.Rooms = roomMgr.RoomManager.GetAllRooms()
	logic.SendMsg(task, cmd.MSG_ListRooms_SC, unsafe.Pointer(ret))
}

func (this *ChatManager) MSG_CreateChatRoom_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_CreateChatRoom_CS) {
	ret := &chat.MSG_CreateChatRoom_SC{}

	//根据房间名创建房间
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

	//加入指定房间
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

	//离开指定房间
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

	//玩家向房间内发送消息
	if !roomMgr.RoomManager.SendMessage(role, msg.Message) {
		ret := &chat.MSG_SendMessage_SC{}
		ret.Message = append(ret.Message, "玩家不在房间或者房间不存在!")
		logic.SendMsg(task, cmd.MSG_SendMessage_SC, unsafe.Pointer(ret))
	}
}

func (this *ChatManager) MSG_SendGMString_CS_CB(task *tcptask.TcpTask, msg *chat.MSG_SendGMString_CS) {
	strArray := strings.Split(msg.Reqgm, " ")
	if len(strArray) < 3 {
		ret := &chat.MSG_SendGMString_SC{Retgm: "GM指令参数非法!"}
		logic.SendMsg(task, cmd.MSG_SendGMString_SC, unsafe.Pointer(ret))
		return
	}

	var retString string
	switch strArray[1] {
	case "/stats":
		roleName := strArray[2]
		roomId := roomMgr.RoomManager.GetRoomIdByRole(roleName)
		loginTime := roleMgr.RoleManager.GetRoleLoginTime(roleName)
		retString = retString + "玩家" + roleName + "信息: "

		if loginTime < 0 {
			retString = retString + "玩家登陆时间非法"
		} else {
			retString = retString + "玩家登陆时间是" + time.Unix(loginTime,0).Format("2006-01-02 15:04:05")
			retString = retString + ",玩家的登陆时长是" + strconv.FormatInt(time.Now().Unix()-loginTime,10)
		}

		if roomId <= 0 {
			retString = retString + ",玩家所在房间非法"
		} else {
			retString = retString + ",玩家所在房间Id为" + strconv.FormatUint(roomId,10)
		}
	case "/popular":
		roomId,_ := strconv.ParseUint(strArray[2],10,64)
		word := roomMgr.RoomManager.GetPopularWord(roomId)
		if word == "" {
			retString = "在房间Id为" + strconv.FormatUint(roomId, 10) + "中没有找到出现频率最高的单词!"
		} else {
			retString = "在房间Id为" + strconv.FormatUint(roomId, 10) + "中出现频率最高的单词是:" + word +  "!"
		}
	}
	ret := &chat.MSG_SendGMString_SC{Retgm: retString }
	logic.SendMsg(task, cmd.MSG_SendGMString_SC, unsafe.Pointer(ret))
}
