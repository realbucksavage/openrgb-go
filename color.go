package openrgb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func readColor(buf []byte, offset int) (Color, error) {
	c := Color{}

	for _, ptr := range []*uint8{&c.Red, &c.Green, &c.Blue} {
		reader := bytes.NewReader(buf[offset:])
		if err := binary.Read(reader, binary.BigEndian, ptr); err != nil {
			return Color{}, err
		}
		offset++
	}

	return c, nil
}

func (c Color) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d);", c.Red, c.Green, c.Blue)
}
