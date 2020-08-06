package openrgb

import (
	"encoding/binary"
	"fmt"
)

type Zone struct {
	Name      string
	Type      uint32
	MinLEDs   uint32
	MaxLEDs   uint32
	TotalLEDs uint32
}

func readZones(buf []byte, count uint16, offset int) ([]Zone, int) {
	zones := make([]Zone, 0)

	for zoneIndex := uint16(0); zoneIndex < count; zoneIndex++ {
		var z Zone
		s, i := readString(buf, offset)
		z.Name = s
		offset += i

		for _, ptr := range []*uint32{
			&z.Type,
			&z.MinLEDs,
			&z.MaxLEDs,
			&z.TotalLEDs,
		} {
			*ptr = binary.LittleEndian.Uint32(buf[offset:])
			offset += 4
		}

		matrixSize := binary.LittleEndian.Uint16(buf[offset:])
		offset += 2 + int(matrixSize)

		zones = append(zones, z)
	}

	return zones, offset
}

func (z Zone) String() string {
	return fmt.Sprintf(`%s (typ %d; LEDs %d)`, z.Name, z.Type, z.TotalLEDs)
}
