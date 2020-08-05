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
	c.sendMessage(commandSetClientName, 0, bytes.NewBufferString("GoClient"))

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

func (c *Client) GetDeviceController(deviceID int) Device {
	c.sendMessage(commandRequestControllerData, deviceID, nil)
	message, _ := c.readMessage()

	var d Device
	d.read(message)

	return d
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

	header := c.decodeHeader(buf)
	buf = make([]byte, header.length)
	_, err = c.clientSock.Read(buf)

	return buf, nil
}

func encodeHeader(command, device, length int) *bytes.Buffer {
	offset := 4
	b := make([]byte, 16)
	for idx, v := range "ORGB" {
		b[idx] = byte(v)
	}

	b[offset] = byte(device)
	offset += 4

	b[offset] = byte(command)
	offset += 4

	b[offset] = byte(length)

	return bytes.NewBuffer(b)
}

func (c *Client) decodeHeader(buffer []byte) orgbHeader {
	return orgbHeader{
		binary.LittleEndian.Uint32(buffer[4:]),
		binary.LittleEndian.Uint32(buffer[8:]),
		binary.LittleEndian.Uint32(buffer[12:]),
	}
}
