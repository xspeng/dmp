package dmpd

import (
	"net"
)

var okBytes = []byte("ok")
var heartbeatBytes = []byte("_heartbeat_")
var seperatorBytes = []byte(" ")

const (
	frameTypeResponse int32 = 0
	frameTypeError    int32 = 1
	frameTypeMessage  int32 = 2
)

type ProtocolV1 struct {
	ctx *context
}

func (p *ProtocolV1) IOLoop(conn net.Conn) error {

	return nil
}

//发送身份验证数据
func (p *ProtocolV1) INDENTIFY() {

}
