//发往socket和从socket接受数据的消息结构
package message

import (
	"bytes"
	"encoding/binary"
	"time"
)

const HEADER_LENGTH = 4
const MAX_PACK_DATA_SIZE = 65535

type Message struct {
	PackSize  uint16 //数据包大小,即不包含PackSize,PackMask的数据大小
	PackMask  uint16 //数据的标志,包括是否压缩,加密等
	MsgId     uint16 //消息Id
	TimeStamp int64  //当前时间戳
	DataSize  uint16 //数据大小,即Data的内容的大小
	Data      []byte //字节流数据
}

func NewMessage(mask uint16, MsgId uint16, DataSize uint16, Data []byte) *Message {
	msg := &Message{DataSize + 12, mask, MsgId, time.Now().Unix(), DataSize, Data}
	return msg
}

//将Message打包成字节流
func (msg *Message) PackBuffer() *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, msg.PackSize)
	binary.Write(buf, binary.LittleEndian, msg.PackMask)
	binary.Write(buf, binary.LittleEndian, msg.MsgId)
	binary.Write(buf, binary.LittleEndian, msg.TimeStamp)
	binary.Write(buf, binary.LittleEndian, msg.DataSize)
	buf.Write(msg.Data)
	return buf
}

//将字节流解包成字节流格式
func (msg *Message) UnpackBuffer(buf *bytes.Buffer) {
	binary.Read(buf, binary.LittleEndian, &msg.PackSize)
	binary.Read(buf, binary.LittleEndian, &msg.PackMask)
	binary.Read(buf, binary.LittleEndian, &msg.MsgId)
	binary.Read(buf, binary.LittleEndian, &msg.TimeStamp)
	binary.Read(buf, binary.LittleEndian, &msg.DataSize)
	msg.Data = make([]byte, buf.Len())
	buf.Read(msg.Data)
}
