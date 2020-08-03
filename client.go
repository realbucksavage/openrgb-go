package openrgb

import (
	"fmt"
	"net"
)

type Client struct {
	clientSock net.Conn
}

func (c *Client) Close() error {
	return c.clientSock.Close()
}

func Connect(host string, port int) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	sock, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{clientSock: sock}, nil
}
