package openrgb

import "encoding/binary"

func (d *Device) read(buf []byte) {

	offset := 4

	d.Type = binary.LittleEndian.Uint32(buf[4:])
	offset += 4

	s, i := readString(buf, offset)
	d.Name = s
	offset += i

	s, i = readString(buf, offset)
	d.Description = s
	offset += i

	s, i = readString(buf, offset)
	d.Version = s
	offset += i

	s, i = readString(buf, offset)
	d.Serial = s
	offset += i

	s, i = readString(buf, offset)
	d.Location = s
	offset += i

	//modeCount := binary.LittleEndian.Uint16(buf[offset:])
	_ = binary.LittleEndian.Uint16(buf[offset:])
	offset += 2

	d.ActiveMode = binary.LittleEndian.Uint32(buf[offset:])
	offset += 4
}

func readString(buf []byte, offset int) (string, int) {
	length := int(binary.LittleEndian.Uint16(buf[offset:]))
	b := buf[offset+2 : offset+length+1]

	return string(b), length + 2
}

func readModes(buf []byte, modeCount uint16, offset int) ([]Mode, int) {
	modes := make([]Mode, 0)

	var modeIndex uint16 = 0
	for ; modeIndex < modeCount; modeIndex++ {
		modeName, i := readString(buf, offset)
		offset += i

		val := binary.LittleEndian.Uint32(buf[offset:])
		offset += 4

		flags := binary.LittleEndian.Uint32(buf[offset:])
		offset += 4

		modes = append(modes, Mode{
			Name:      modeName,
			Value:     val,
			Flags:     flags,
			MinSpeed:  0,
			MaxSpeed:  0,
			MinColors: 0,
			MaxColors: 0,
			Speed:     0,
			Direction: 0,
			ColorMode: 0,
			Colors:    nil,
		})
	}

	return modes, 0
}
