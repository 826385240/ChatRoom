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
	fmt.Println("请输入创建/登陆的用户名")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ret := &login.MSG_CreateRole_CS{Name: input.Text()}
	logic.SendMsg(client, cmd.MSG_CreateRole_CS, unsafe.Pointer(ret))
}

func (this *LoginManager) MSG_CreateRole_SC_CB(client *tcpclient.TcpClient, msg *login.MSG_CreateRole_SC) {
	if msg.Retcode {
		fmt.Println("**************************创建成功!***************************")
	} else {
		fmt.Println("*********************角色已存在,登陆成功!***********************")
	}

	fmt.Println("=============================================================")
	fmt.Println("请输入接下来的操作:")
	fmt.Println("1.listrooms ==> 列出所有聊天室")
	fmt.Println("2.createroom 房间名 ==> 创建指定房间")
	fmt.Println("3.joinroom 房间名 ==> 加入指定房间")
	fmt.Println("4.leaveroom ==> 离开当前房间")
	fmt.Println("5.sendmsg 聊天消息 ==> 发送聊天消息")
	fmt.Println("6.sendgm GM命令 ==> 发送GM指令")
	fmt.Println("=============================================================")

	operateMgr.OperateMgr.ShowOperate(client)
}

func StartLogin(client *tcpclient.TcpClient) {
	ret := &login.MSG_StartLogin_CS{}
	logic.SendMsg(client, cmd.MSG_StartLogin_CS, unsafe.Pointer(ret))
}

