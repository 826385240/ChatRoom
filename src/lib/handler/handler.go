package handler

import (
	"ChatRoom/src/lib/logger"
	"unsafe"
)

type ICbHandle interface {
	Exec(unsafe.Pointer, unsafe.Pointer)
}

type Handler struct {
	handleTable map[uint16]ICbHandle
}

func NewHandler() *Handler {
	return &Handler{make(map[uint16]ICbHandle)}
}

func (this *Handler) IsExist(msgId uint16) bool {
	_, ok := this.handleTable[msgId]
	return ok
}

func (this *Handler) GetHandler(msgId uint16) ICbHandle {
	h := this.handleTable[msgId]
	return h
}

func (this *Handler) RegHandler(msgId uint16, cb ICbHandle) bool {
	_, ok := this.handleTable[msgId]
	if !ok {
		this.handleTable[msgId] = cb
		return true
	}

	return false
}

func (this *Handler) ExecHandler(msgId uint16, u unsafe.Pointer, m unsafe.Pointer) bool {
	h, ok := this.handleTable[msgId]
	if ok {
		h.Exec(u, m)
		return true
	}
	logger.DEBUG("错误!%d消息找不了处理函数!", msgId)
	return false
}

type ICmdHandle interface {
	Init()
}
