package roomMgr

import (
	cmd "ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/serverlogic/roleMgr"
	"container/list"
	"unsafe"
)

type Room struct {
	roomRoles map[string]bool
	allMessages *list.List
}

type roomManager struct {
	allRooms map[string]Room
	role2room map[string]string
}

var RoomManager = &roomManager{allRooms: map[string]Room{}, role2room: map[string]string{}}

func (this *roomManager) CreateRoom(n string) bool {
	_,ok := this.allRooms[n]

	if !ok {
		this.allRooms[n]=Room{roomRoles: map[string]bool{}, allMessages: list.New()}
	}
	return !ok
}

func (this *roomManager) JoinRoom(task *tcptask.TcpTask, role string, room string) bool {
	_,ok := this.allRooms[room]
	if !ok {
		return ok
	}

	r := this.allRooms[room]
	_,ok1 := r.roomRoles[role]
	if ok1 {
		return false
	}

	r.roomRoles[role]=true
	this.role2room[role]=room

	ret1 := &chat.MSG_JoinRoom_SC{Retcode: true}
	logic.SendMsg(task, cmd.MSG_JoinRoom_SC, unsafe.Pointer(ret1))

	//进入房间广播50条消息
	ret2 := &chat.MSG_SendMessage_SC{}
	var i int = 0
	for e := r.allMessages.Front(); e!=nil; e=e.Next() {
		i=i+1
		if i > 50 {
			break
		}
		ret2.Message = append(ret2.Message, (e.Value).(string))
	}
	logic.SendMsg(task, cmd.MSG_SendMessage_SC, unsafe.Pointer(ret2))
	return true
}

func (this *roomManager) LeaveRoom(role string) bool {
	room,ok1 := this.role2room[role]
	if !ok1 {
		return ok1
	}

	_,ok2 := this.allRooms[room]
	if !ok2 {
		return ok2
	}

	r := this.allRooms[room]
	_,ok3 := r.roomRoles[role]
	if !ok3 {
		return ok3
	}

	delete(r.roomRoles, role)
	delete(this.role2room, role)

	if len(r.roomRoles) <= 0 {
		delete(this.allRooms, room)
	}
	return true
}

func (this *roomManager) SendMessage(role string, msg string) bool {
	room,ok1 := this.role2room[role]
	if !ok1 {
		return ok1
	}

	_,ok2 := this.allRooms[room]
	if !ok2 {
		return ok2
	}

	r := this.allRooms[room]
	_,ok3 := r.roomRoles[role]
	if !ok3 {
		return ok3
	}

	r.allMessages.PushBack(role + ":" + msg)

	for k,_ := range r.roomRoles {
		task := roleMgr.RoleManager.GetTcpTask(k)
		if task != nil {
			ret := &chat.MSG_SendMessage_SC{}
			ret.Message = append(ret.Message, (r.allMessages.Back().Value).(string))
			logic.SendMsg(task, cmd.MSG_SendMessage_SC, unsafe.Pointer(ret))
		}
	}
	return true
}

func (this *roomManager) GetAllRooms() []string  {
	ret := make([]string,0,1)
	for k,_ := range this.allRooms{
		ret=append(ret, k)
	}
	return ret
}
