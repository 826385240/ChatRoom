package ipaddr

import (
	"net"
	"strconv"
)

type IpAddr struct {
	Proto string
	Ip    string
	Port  uint16
}

func (this *IpAddr) Listen() (net.Listener, error) {
	if net.ParseIP(this.Ip) == nil {
		return nil, nil
	}

	strPort := strconv.Itoa(int(this.Port))
	addr := this.Ip + ":" + strPort
	return net.Listen(this.Proto, addr)
}

func (this *IpAddr) Connect() (net.Conn, error) {
	if net.ParseIP(this.Ip) == nil {
		return nil, nil
	}

	strPort := strconv.Itoa(int(this.Port))
	addr := this.Ip + ":" + strPort
	return net.Dial(this.Proto, addr)
}
