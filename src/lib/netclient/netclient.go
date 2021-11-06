package netclient

import (
	"net"
	"strconv"
)

type BaseNetClient struct {
}

func NewBaseNetClient() *BaseNetClient {
	return &BaseNetClient{}
}

func (this *BaseNetClient) Connect(netProto string, ip string, port uint16) (net.Conn, error) {
	strPort := strconv.Itoa(int(port))
	addr := ip + ":" + strPort
	return net.Dial(netProto, addr)
}
