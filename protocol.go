package protocol

import (
	"fmt"
	"net"
)

type TCPHandler interface {
	Handle(net.Conn)
}

func TCPServer(listener net.Listener, handler TCPHandler) {
	fmt.Println("TCP: server start ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("listen on address %s failed", listener.Addr())
			break
		}
		go handler.Handle(conn)
	}
	fmt.Println("TCP server shutdown")
}
