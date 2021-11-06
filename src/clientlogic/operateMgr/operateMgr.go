package operateMgr

import (
	"ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcpclient"
	"ChatRoom/src/protoout/chat"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

type operateManager struct {
}

var OperateMgr = &operateManager{}

func (this *operateManager) ShowOperate(client* tcpclient.TcpClient){
	fmt.Println("=============================================================")
	fmt.Println("请输入接下来的操作:")
	fmt.Println("1.listrooms ==> 列出所有聊天室")
	fmt.Println("2.createroom 房间名 ==> 创建指定房间")
	fmt.Println("3.joinroom 房间名 ==> 加入指定房间")
	fmt.Println("4.leaveroom ==> 离开当前房间")
	fmt.Println("5.sendmsg 聊天消息 ==> 发送聊天消息")
	fmt.Println("6.sendgm GM命令 ==> 发送GM指令")
	fmt.Println("=============================================================")

	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		strArray := strings.Split(input.Text(), " ")
		if len(strArray) <= 0 {
			fmt.Println("参数错误,请重新输入!")
		}

		// 根据玩家输入执行相应操作
		switch strArray[0] {
		case "listrooms":
			this.ListRooms(client)
			return;
		case "createroom":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.CreateRoom(client, strArray[1])
				return;
			}
		case "joinroom":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.JoinRoom(client, strArray[1])
				return;
			}
		case "leaveroom":
			this.LeaveRoom(client)
			return;
		case "sendmsg":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.SendMsg(client, strArray[1])
				return;
			}
		case "sendgm":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.SendGM(client, strArray[1])
				return;
			}
		default:
			fmt.Println("参数错误,请重新输入!")
		}
	}
}

func (this *operateManager) ListRooms(client *tcpclient.TcpClient) {
	ret := &chat.MSG_ListRooms_CS{}
	logic.SendMsg(client, cmd.MSG_ListRooms_CS, unsafe.Pointer(ret))
}

func (this *operateManager) CreateRoom(client *tcpclient.TcpClient, n string) {
	ret := &chat.MSG_CreateChatRoom_CS{Name: n}
	logic.SendMsg(client, cmd.MSG_CreateChatRoom_CS, unsafe.Pointer(ret))
}

func (this *operateManager) JoinRoom(client *tcpclient.TcpClient, n string) {
	ret := &chat.MSG_JoinRoom_CS{Name: n}
	logic.SendMsg(client, cmd.MSG_JoinRoom_CS, unsafe.Pointer(ret))
}

func (this *operateManager) LeaveRoom(client *tcpclient.TcpClient) {
	ret := &chat.MSG_LeaveRoom_CS{}
	logic.SendMsg(client, cmd.MSG_LeaveRoom_CS, unsafe.Pointer(ret))
}

func (this *operateManager) SendMsg(client *tcpclient.TcpClient, m string) {
	ret := &chat.MSG_SendMessage_CS{Message: m}
	logic.SendMsg(client, cmd.MSG_SendMessage_CS, unsafe.Pointer(ret))
}

func (this *operateManager) SendGM(client *tcpclient.TcpClient, m string) {
	ret := &chat.MSG_SendGMString_CS{Reqgm: m}
	logic.SendMsg(client, cmd.MSG_SendGMString_CS, unsafe.Pointer(ret))
}
