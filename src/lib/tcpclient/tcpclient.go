//tcpclient主要提供协程间通信的接口并向socket转发接收数据
package tcpclient

import (
	"ChatRoom/src/lib/common"
	"ChatRoom/src/lib/logger"
	"ChatRoom/src/lib/message"
	"ChatRoom/src/lib/socket"
	"bytes"
	"encoding/binary"
	"net"
	"sync"
	"unsafe"
)

type TcpClient struct {
	uniqId uint64   //每个连接的临时的唯一标识符
	conn   net.Conn //task所管理的连接
	*socket.Socket
	w_chan chan com.MsgToLogicPtr //task向逻辑协程写入数据的channel
	r_chan chan com.MsgToLogicPtr //task从逻辑协程读取数据的channel
	flag   uint16
	wg     sync.WaitGroup
}

func NewTcpClient(conn net.Conn) *TcpClient {
	var sock *socket.Socket = socket.NewSocket(conn)
	rc := make(chan com.MsgToLogicPtr, com.READ_CHAN_BUFFER_SIZE)
	client := &TcpClient{conn: conn, Socket: sock, w_chan: nil, r_chan: rc}
	client.uniqId = *(*uint64)(unsafe.Pointer(&client))
	return client
}

func (this *TcpClient) WgAdd(i int) {
	this.wg.Add(i)
}

func (this *TcpClient) WgWait() {
	this.wg.Wait()
}

func (this *TcpClient) WgDone() {
	this.wg.Done()
}

func (this *TcpClient) GetUniqId() uint64 {
	return this.uniqId
}

func (this *TcpClient) SetWChan(ch chan com.MsgToLogicPtr) {
	this.w_chan = ch
}

func (this *TcpClient) GetTaskFlag() uint16 {
	return this.flag
}

func (this *TcpClient) SetTaskFlag(i uint16) {
	this.flag = this.flag | i
}

func (this *TcpClient) SendCmd(msgId uint16, data []byte) bool {
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

func (this *TcpClient) RecvCmd() *message.Message {
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
func (this *TcpClient) RecvFromRChanNB() com.MsgToLogicPtr {
	select {
	case d := <-this.r_chan:
		return d
	default:
		return nil
	}
}

//非阻塞向r_chan写入一个发送对象
func (this *TcpClient) SendToRChanNB(d com.MsgToLogicPtr) bool {
	select {
	case this.r_chan <- d:
		return true
	default:
		return false
	}
}

//非阻塞从w_chan读取一个发送对象
func (this *TcpClient) RecvFromWChanNB() com.MsgToLogicPtr {
	select {
	case d := <-this.w_chan:
		return d
	default:
		return nil
	}
}

//非阻塞向w_chan写入一个发送对象
func (this *TcpClient) SendToWChanNB(d com.MsgToLogicPtr) bool {
	select {
	case this.w_chan <- d:
		return true
	default:
		return false
	}
}

//阻塞从r_chan读取一个发送对象
func (this *TcpClient) RecvFromRChan() com.MsgToLogicPtr {
	d := <-this.r_chan
	return d
}

//阻塞向r_chan写入一个发送对象
func (this *TcpClient) SendToRChan(d com.MsgToLogicPtr) {
	this.r_chan <- d
}

//阻塞从w_chan读取一个发送对象
func (this *TcpClient) RecvFromWChan() com.MsgToLogicPtr {
	d := <-this.w_chan
	return d
}

//阻塞向w_chan写入一个发送对象
func (this *TcpClient) SendToWChan(d com.MsgToLogicPtr) {
	this.w_chan <- d
}

//关闭写协程
func (this *TcpClient) CloseWChan() {
	close(this.w_chan)
}

//关闭读协程
func (this *TcpClient) CloseRChan() {
	close(this.r_chan)
}
