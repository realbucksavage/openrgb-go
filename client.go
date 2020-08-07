package openrgb

import (
	"bytes"
	"encoding/binary"
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

	c := &Client{clientSock: sock}

	err = c.sendMessage(commandSetClientName, 0, bytes.NewBufferString("GoClient"))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) GetControllerCount() (int, error) {
	err := c.sendMessage(commandRequestControllerCount, 0, nil)
	if err != nil {
		return 0, err
	}

	message, _ := c.readMessage()
	count := int(binary.LittleEndian.Uint32(message))

	return count, nil
}

func (c *Client) GetDeviceController(deviceID int) (Device, error) {
	if err := c.sendMessage(commandRequestControllerData, deviceID, nil); err != nil {
		return Device{}, err
	}
	message, _ := c.readMessage()

	d, err := readDevice(message)
	if err != nil {
		return Device{}, err
	}

	return d, nil
}

func (c *Client) sendMessage(command, deviceID int, buffer *bytes.Buffer) error {
	bufLen := 0
	if buffer != nil {
		bufLen = buffer.Len()
	}

	header := encodeHeader(command, deviceID, bufLen)
	if buffer != nil {
		header.Write(buffer.Bytes())
	}

	_, err := c.clientSock.Write(header.Bytes())
	if err != nil {
		return err
	}

	return err
}

func (c *Client) readMessage() ([]byte, error) {
	buf := make([]byte, 16)
	_, err := c.clientSock.Read(buf)
	if err != nil {
		return nil, err
	}

	header := decodeHeader(buf)
	buf = make([]byte, header.length)
	_, err = c.clientSock.Read(buf)

	return buf, nil
}
