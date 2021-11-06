package loginMgr

import (
	"ChatRoom/src/clientlogic/operateMgr"
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcpclient"
	"ChatRoom/src/protoout/login"
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

type LoginManager struct {
}

func (this *LoginManager) Init() {
	cmd.Reg_TcpClient_MSG_StartLogin_SC(this.MSG_StartLogin_SC_CB)
	cmd.Reg_TcpClient_MSG_CreateRole_SC(this.MSG_CreateRole_SC_CB)
}

func (this *LoginManager) MSG_StartLogin_SC_CB(client *tcpclient.TcpClient, msg *login.MSG_StartLogin_SC) {
	fmt.Println("请输入创建的用户名")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ret := &login.MSG_CreateRole_CS{Name: input.Text()}
	logic.SendMsg(client, cmd.MSG_CreateRole_CS, unsafe.Pointer(ret))
}

func (this *LoginManager) MSG_CreateRole_SC_CB(client *tcpclient.TcpClient, msg *login.MSG_CreateRole_SC) {
	if msg.Retcode {
		fmt.Println("登陆成功!")
	}

	operateMgr.OperateMgr.ShowOperate(client)
}

func StartLogin(client *tcpclient.TcpClient) {
	ret := &login.MSG_StartLogin_CS{}
	logic.SendMsg(client, cmd.MSG_StartLogin_CS, unsafe.Pointer(ret))
}

