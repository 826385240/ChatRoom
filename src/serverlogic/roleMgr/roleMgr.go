package roleMgr

import (
	"ChatRoom/src/lib/tcptask"
	"time"
)

type Role struct{
	loginTime int64
	name string
	task *tcptask.TcpTask
}

type roleManager struct {
	allRoles map[string]Role
	client2role map[*tcptask.TcpTask]string
}

var RoleManager=&roleManager{allRoles: map[string]Role{}, client2role: map[*tcptask.TcpTask]string{}}

func (this *roleManager) CreateRole(t *tcptask.TcpTask, n string) bool {
	_,ok := this.allRoles[n]
	this.client2role[t]=n
	this.allRoles[n]=Role{name: n, task:t, loginTime:time.Now().Unix()}
	return !ok
}

func (this *roleManager) GetRoleName(task *tcptask.TcpTask) string {
	n,ok := this.client2role[task]

	if ok {
		return n
	}
	return ""
}

func (this *roleManager) GetTcpTask(name string) *tcptask.TcpTask {
	r,ok := this.allRoles[name]
	if !ok {
		return nil
	}

	return r.task
}
func (this *roleManager) GetRoleLoginTime(name string) int64 {
	r,ok := this.allRoles[name]
	if !ok {
		return -1
	}

	return r.loginTime
}
