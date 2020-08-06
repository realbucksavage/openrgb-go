package openrgb

import (
	"bytes"
	"encoding/binary"
)

type orgbHeader struct {
	deviceID  uint32
	commandID uint32
	length    uint32
}

func readString(buf []byte, offset int) (string, int) {
	length := int(binary.LittleEndian.Uint16(buf[offset:]))
	b := buf[offset+2 : offset+length+1]

	return string(b), length + 2
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

func decodeHeader(buffer []byte) orgbHeader {
	return orgbHeader{
		binary.LittleEndian.Uint32(buffer[4:]),
		binary.LittleEndian.Uint32(buffer[8:]),
		binary.LittleEndian.Uint32(buffer[12:]),
	}
}
