package main

import (
	"ChatRoom/src/clientlogic/chatMgr"
	"ChatRoom/src/clientlogic/loginMgr"
	"ChatRoom/src/lib/client"
	"ChatRoom/src/lib/logger"
	"runtime"
)

func OnTimer(tm int64) {

}

func main() {
	runtime.GOMAXPROCS(2)
	logger.InitLog("./log/mainclient.log")

	c := client.NewClient()
	tempclient := c.Connect("tcp", "127.0.0.1", 8000)
	if tempclient == nil {
		logger.FATAL("错误!连接服务器失败!")
	}

	c.AddCmdHandle(&loginMgr.LoginManager{})
	c.AddCmdHandle(&chatMgr.ChatManager{})

	loginMgr.StartLogin(tempclient)
	c.Start(OnTimer)
}
