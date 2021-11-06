package exception

import (
	"ChatRoom/src/lib/logger"
)

func ErrorExit(err error) {
	if err != nil {
		logger.PANIC("error:%s", err.Error())
	}
}

func StringExit(err string) {
	logger.PANIC("error:%s", err)
}
