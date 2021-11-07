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
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		inputString := input.Text()
		strArray := strings.Split(inputString, " ")
		if len(strArray) <= 0 {
			fmt.Println("参数错误,请重新输入!")
		}

		// 根据玩家输入执行相应操作
		switch strArray[0] {
		case "listrooms":
			this.ListRooms(client)
			return
		case "createroom":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.CreateRoom(client, strArray[1])
				return
			}
		case "joinroom":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.JoinRoom(client, strArray[1])
				return
			}
		case "leaveroom":
			this.LeaveRoom(client)
			return
		case "sendmsg":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				inputString = strings.Replace(inputString, "sendmsg ", "", 1)
				this.SendMsg(client, inputString)
				return
			}
		case "sendgm":
			if len(strArray) <2 {
				fmt.Println("参数错误,请重新输入!")
			} else {
				this.SendGM(client, inputString)
				return
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
