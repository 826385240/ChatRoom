//tcptask主要提供协程间通信的接口并向socket转发接收数据
package tcptask

import (
	"bytes"
	"encoding/binary"
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/logger"
	"ChatRoom/src/lib/message"
	"ChatRoom/src/lib/socket"
	"net"
	"sync"
	"unsafe"
)

type TcpTask struct {
	uniqId uint64   //每个连接的临时的唯一标识符
	conn   net.Conn //task所管理的连接
	*socket.Socket
	w_chan chan com.MsgToLogicPtr //task向逻辑协程写入数据的channel
	r_chan chan com.MsgToLogicPtr //task从逻辑协程读取数据的channel
	flag   uint16
	wg     sync.WaitGroup
}

func NewTcpTask(conn net.Conn) *TcpTask {
	var sock *socket.Socket = socket.NewSocket(conn)
	rc := make(chan com.MsgToLogicPtr, com.READ_CHAN_BUFFER_SIZE)
	task := &TcpTask{conn: conn, Socket: sock, w_chan: nil, r_chan: rc}
	task.uniqId = *(*uint64)(unsafe.Pointer(&task))
	return task
}

func (this *TcpTask) WgAdd(i int) {
	this.wg.Add(i)
}

func (this *TcpTask) WgWait() {
	this.wg.Wait()
}

func (this *TcpTask) WgDone() {
	this.wg.Done()
}

func (this *TcpTask) GetUniqId() uint64 {
	return this.uniqId
}

func (this *TcpTask) SetWChan(ch chan com.MsgToLogicPtr) {
	this.w_chan = ch
}

func (this *TcpTask) GetTaskFlag() uint16 {
	return this.flag
}

func (this *TcpTask) SetTaskFlag(i uint16) {
	this.flag = this.flag | i
}

func (this *TcpTask) SendCmd(msgId uint16, data []byte) bool {
	length := len(data)
	if length >= (message.MAX_PACK_DATA_SIZE - message.HEADER_LENGTH) {
		logger.ERROR("错误!数据包长度过大!消息ID:%d", msgId)
		return false
	}

	//将发送的数据进行打包
	msg := message.NewMessage(this.GetTaskFlag(), msgId, uint16(length), data)
	buf := msg.PackBuffer()

	this.WriteDataToBuffer(buf.Bytes())
	for {
		succ, err := this.WriteToConn()
		if err != nil {
			//发送错误,打印错误退出循环
			logger.ERROR("错误!发送数据错误!%ss", err.Error())
			return false
		} else if succ {
			//发送完毕,直接退出循环
			return true
		}
	}
}

func (this *TcpTask) RecvCmd() *message.Message {
	for {
		//至少需要读取一个完整消息才会退出
		curLength := this.ReadBufferSize()
		if curLength > message.HEADER_LENGTH {
			data := this.GetReadBuffer()
			dataLength := binary.LittleEndian.Uint16(data[:])

			if curLength >= (int(dataLength) + message.HEADER_LENGTH) {
				//将接受的数据进行解包
				msg := &message.Message{}
				msg.UnpackBuffer(bytes.NewBuffer(data[:dataLength+message.HEADER_LENGTH]))
				this.ReadDone(int(dataLength + message.HEADER_LENGTH))
				return msg
			}
		}

		//读取数据失败,关闭连接
		if this.ReadFromConn() != nil {
			return nil
		}
	}
}

//非阻塞从r_chan读取一个发送对象
func (this *TcpTask) RecvFromRChanNB() com.MsgToLogicPtr {
	select {
	case d := <-this.r_chan:
		return d
	default:
		return nil
	}
}

//非阻塞向r_chan写入一个发送对象
func (this *TcpTask) SendToRChanNB(d com.MsgToLogicPtr) bool {
	select {
	case this.r_chan <- d:
		return true
	default:
		return false
	}
}

//非阻塞从w_chan读取一个发送对象
func (this *TcpTask) RecvFromWChanNB() com.MsgToLogicPtr {
	select {
	case d := <-this.w_chan:
		return d
	default:
		return nil
	}
}

//非阻塞向w_chan写入一个发送对象
func (this *TcpTask) SendToWChanNB(d com.MsgToLogicPtr) bool {
	select {
	case this.w_chan <- d:
		return true
	default:
		return false
	}
}

//阻塞从r_chan读取一个发送对象
func (this *TcpTask) RecvFromRChan() com.MsgToLogicPtr {
	d := <-this.r_chan
	return d
}

//阻塞向r_chan写入一个发送对象
func (this *TcpTask) SendToRChan(d com.MsgToLogicPtr) {
	this.r_chan <- d
}

//阻塞从w_chan读取一个发送对象
func (this *TcpTask) RecvFromWChan() com.MsgToLogicPtr {
	d := <-this.w_chan
	return d
}

//阻塞向w_chan写入一个发送对象
func (this *TcpTask) SendToWChan(d com.MsgToLogicPtr) {
	this.w_chan <- d
}

//关闭写协程
func (this *TcpTask) CloseWChan() {
	close(this.w_chan)
}

//关闭读协程
func (this *TcpTask) CloseRChan() {
	close(this.r_chan)
}
