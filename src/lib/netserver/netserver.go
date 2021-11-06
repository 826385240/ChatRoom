/*网络服务器服务器的功能:
 *1.记录ip和端口信息,记录分区信息
 *2.绑定和监听端口
 *3.接受新的连接
 */
package netserver

import (
	"ChatRoom/src/lib/ipaddr"
	"net"
	"strconv"
)

type ZoneType struct {
	zoneType uint16
	zoneId   uint16
}

type BaseNetServer struct {
	innerAddr ipaddr.IpAddr
	outerAddr ipaddr.IpAddr
	zoneInfo  ZoneType
}

func NewBaseNetServer(innerProto string, innerIp string, innerPort uint16, outerProto string, outerIp string, outerPort uint16) *BaseNetServer {
	iAddr := ipaddr.IpAddr{innerProto, innerIp, innerPort}
	oAddr := ipaddr.IpAddr{outerProto, outerIp, outerPort}
	return &BaseNetServer{innerAddr: iAddr, outerAddr: oAddr}
}

func (this *BaseNetServer) Listen(netProto string, ip string, port uint16) (net.Listener, error) {
	strPort := strconv.Itoa(int(port))
	addr := ip + ":" + strPort
	return net.Listen(netProto, addr)
}

func (this *BaseNetServer) Accept(liconn net.Listener) (net.Conn, error) {
	return liconn.Accept()
}

func (this *BaseNetServer) InnerListen() (net.Listener, error) {
	return this.innerAddr.Listen()
}

func (this *BaseNetServer) OuterListen() (net.Listener, error) {
	return this.outerAddr.Listen()
}

func (this *BaseNetServer) getInnerProto() string {
	return this.innerAddr.Proto
}

func (this *BaseNetServer) getInnerIp() string {
	return this.innerAddr.Ip
}

func (this *BaseNetServer) getInnerPort() uint16 {
	return this.innerAddr.Port
}

func (this *BaseNetServer) getOuterProto() string {
	return this.outerAddr.Proto
}

func (this *BaseNetServer) getOuterIp() string {
	return this.outerAddr.Ip
}

func (this *BaseNetServer) getOuterPort() uint16 {
	return this.outerAddr.Port
}

func (this *BaseNetServer) getZoneType() uint16 {
	return this.zoneInfo.zoneType
}

func (this *BaseNetServer) getZoneId() uint16 {
	return this.zoneInfo.zoneId
}
