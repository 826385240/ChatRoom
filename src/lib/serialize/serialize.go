//protobuf序列化与反序列化
package serialize

import "github.com/golang/protobuf/proto"

func SerializeToBuffer(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func UnserializeToBuffer(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}
