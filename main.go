package main

import "u2t/u2t"

const (
	ip   = "0.0.0.0"
	port = 5001
)

func main() {
	go u2t.TCPServer("127.0.0.1", 12345)
	// go t2u.UDPServer(ip, 1399)

	// go u2t.UdpToTcpMain(ip, port)
	// t2u.TcpToUdpMain(ip, 8090)
	u2t.UdpToTcpMain(ip, port)
}
