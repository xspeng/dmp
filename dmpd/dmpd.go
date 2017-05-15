package dmpd

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/xspeng/dmp/protocol"
	"github.com/xspeng/dmp/util"
)

//作为服务端进程
type DMPD struct {
	ctx *context

	TCPAddress   string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	net.Conn
	tcpListener net.Listener
	reader      bufio.Reader
	writer      bufio.Writer

	sync.RWMutex
	wg util.WaitGroupWrapper

	topicMap map[string]*Topic
}

func (d *DMPD) Main() {
	listener, err := net.Listen("tcp", d.TCPAddress)
	if err != nil {
		fmt.Errorf("failed listen on address :%s", d.TCPAddress)
	}
	d.Lock()
	d.tcpListener = listener
	d.Unlock()
	ctx := &context{d}
	server := &tcpServer{ctx}
	d.wg.Wrap(func() {
		protocol.TCPServer(listener, server)
	})
	d.wg.Wait()
}

func New(opts *Options) *DMPD {
	d := &DMPD{
		TCPAddress:   opts.TCPAddress,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
	}
	d.ctx = &context{d}
	return d
}
