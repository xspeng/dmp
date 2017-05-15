package dmpd

import (
	"net"
)

type Client struct {
	ID string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Handle(conn net.Conn) {

}
