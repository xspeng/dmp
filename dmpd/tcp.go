package dmpd

import (
	"fmt"
	"io"
	"net"

	"github.com/xspeng/dmp/protocol"
)

type tcpServer struct {
	ctx *context
}

func (p *tcpServer) Handle(conn net.Conn) {
	fmt.Printf("TCP: new client from address:%s\n", conn.RemoteAddr())

	//var buf [4]byte
	buf := make([]byte, 4)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Errorf("failed when read :%s\n", err)
	}
	protocolMagic := string(buf)
	var prot protocol.Protocol
	switch protocolMagic {
	case "  V1":
		prot = &ProtocolV1{p.ctx}
		fmt.Println("use protocol magic V1")
		if err != nil {

		}
	default:
		fmt.Errorf("unknown protocol version:%s\n ", protocolMagic)
	}
	err = prot.IOLoop(conn)
	if err != nil {
		fmt.Errorf("", err)
	}
	return
}
