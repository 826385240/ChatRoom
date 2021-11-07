//socket主要负责二进制数据的发送与接收
package socket

import (
	"ChatRoom/src/lib/logger"
	"io"
	"net"
)

const RECV_BUFFER_SIZE = 8092

type Socket struct {
	tempBuf     []byte
	conn        net.Conn
	ReadBuffer  []byte
	WriteBuffer []byte
}

func NewSocket(conn net.Conn) *Socket {
	sock := &Socket{conn: conn, tempBuf: make([]byte, RECV_BUFFER_SIZE, RECV_BUFFER_SIZE), ReadBuffer: []byte{}, WriteBuffer: []byte{}}
	return sock
}

func (this *Socket) ReadFromConn() error {
	//for {
	n, err := this.conn.Read(this.tempBuf[0:])
	if n > 0 {
		this.ReadBuffer = append(this.ReadBuffer, this.tempBuf[0:n]...)
		if err != nil {
			if err != io.EOF {
				logger.ERROR("错误!从Socket读取数据失败!%s", err.Error())
			}
			return err
		}
	} else {
		return err
	}

	//}
	return nil
}

func (this *Socket) WriteDataToBuffer(data []byte) {
	this.WriteBuffer = append(this.WriteBuffer, data[0:]...)
}

/*
 * 发送一次数据
 * 返回值: 是否发送完毕,error
 */
func (this *Socket) WriteToConn() (bool, error) {
	n, err := this.conn.Write(this.WriteBuffer)
	if err != nil {
		return false, err
	}

	//数据发送完毕
	if n >= len(this.WriteBuffer) {
		this.WriteBuffer = this.WriteBuffer[:0]
		return true, nil
	}

	//数据没有发送完毕
	this.WriteBuffer = this.WriteBuffer[n:]
	return false, nil
}

func (this *Socket) ReadBufferSize() int {
	return len(this.ReadBuffer)
}

func (this *Socket) WriteBufferSize() int {
	return len(this.WriteBuffer)
}

func (this *Socket) GetReadBuffer() []byte {
	return this.ReadBuffer[:]
}

func (this *Socket) GetWriteBuffer() []byte {
	return this.ReadBuffer[:]
}

func (this *Socket) ReadDone(size int) {
	this.ReadBuffer = this.ReadBuffer[size:]
}

func (this *Socket) WriteDone(size int) {
	this.WriteBuffer = this.WriteBuffer[size:]
}

func (this *Socket) CloseConn() {
	this.conn.Close()
}
