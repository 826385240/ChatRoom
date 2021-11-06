package loginMgr

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/login"
	"ChatRoom/src/serverlogic/roleMgr"
	"unsafe"
)

type LoginManager struct {
}

func (this *LoginManager) Init() {
	cmd.Reg_TcpTask_MSG_StartLogin_CS(this.MSG_StartLogin_CS_CB)
	cmd.Reg_TcpTask_MSG_CreateRole_CS(this.MSG_CreateRole_CS_CB)
}

func (this *LoginManager) MSG_StartLogin_CS_CB(task *tcptask.TcpTask, msg *login.MSG_StartLogin_CS) {
	ret := &login.MSG_StartLogin_SC{}
	logic.SendMsg(task, cmd.MSG_StartLogin_SC, unsafe.Pointer(ret))
}

func (this *LoginManager) MSG_CreateRole_CS_CB(task *tcptask.TcpTask, msg *login.MSG_CreateRole_CS) {
	//如果不存在在就创建角色
	roleMgr.RoleManager.CreateRole(task, msg.GetName())

	ret := &login.MSG_CreateRole_SC{Retcode: true}
	logic.SendMsg(task, cmd.MSG_CreateRole_SC, unsafe.Pointer(ret))
}

