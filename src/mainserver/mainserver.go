package main

import (
	"ChatRoom/src/lib/benchmark"
	"ChatRoom/src/lib/exception"
	"ChatRoom/src/lib/logger"
	"ChatRoom/src/lib/server"
	"ChatRoom/src/serverlogic/chatMgr"
	"ChatRoom/src/serverlogic/loginMgr"
	"runtime"
)

func Ontimer(tm int64) {

}

func main() {
	runtime.GOMAXPROCS(2)
	logger.InitLog("./log/mainserver.log")
	//初始化服务器对象
	server := server.NewServer("tcp", "127.0.0.1", 9000, "tcp", "0.0.0.0", 8000)
	if server == nil {
		exception.StringExit("错误!启动服务器失败")
	}

	//性能测试
	benchmark.LocalProfile()
	defer benchmark.DeferLocalProfile()

	server.AddCmdHandle(&loginMgr.LoginManager{})
	server.AddCmdHandle(&chatMgr.ChatManager{})
	server.Start(Ontimer)
}
