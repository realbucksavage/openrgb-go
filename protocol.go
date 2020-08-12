package openrgb

import (
	"bytes"
	"encoding/binary"
)

var (
	offset8BEBits  = 1
	offset16LEBits = 2
	offset32LEBits = 4
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

func encodeHeader(header orgbHeader) *bytes.Buffer {
	b := bytes.NewBufferString("ORGB")

	for _, v := range []uint32{
		header.deviceID,
		header.commandID,
		header.length,
	} {
		buf := make([]byte, offset32LEBits)
		binary.LittleEndian.PutUint32(buf, v)
		b.Write(buf)
	}

	return b
}

func decodeHeader(buffer []byte) orgbHeader {
	return orgbHeader{
		binary.LittleEndian.Uint32(buffer[4:]),
		binary.LittleEndian.Uint32(buffer[8:]),
		binary.LittleEndian.Uint32(buffer[12:]),
	}
}
