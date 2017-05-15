package protocol

import (
	"fmt"
	"net"
)

type TCPHandler interface {
	Handle(net.Conn)
}

func TCPServer(listener net.Listener, handler TCPHandler) {
	fmt.Println("TCP: server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("failed accept on address :%s", listener.Addr())
			continue
		}
		go handler.Handle(conn)
	}
	fmt.Println("TCP: server shutdown")
}
