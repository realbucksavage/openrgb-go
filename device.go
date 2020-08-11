package openrgb

import (
	"encoding/binary"
	"fmt"
)

type Device struct {
	Type        uint32
	Name        string
	Description string
	Version     string
	Serial      string
	Location    string
	ActiveMode  uint32
	LEDs        []LED
	Colors      []Color
	Modes       []Mode
	Zones       []Zone
}

func (d *Device) read(buf []byte) error {

	offset := offset32LEBits

	d.Type = binary.LittleEndian.Uint32(buf[4:])
	offset += offset32LEBits

	for _, st := range []*string{
		&d.Name,
		&d.Description,
		&d.Version,
		&d.Serial,
		&d.Location,
	} {
		s, i := readString(buf, offset)
		offset += i
		*st = s
	}

	modeCount := binary.LittleEndian.Uint16(buf[offset:])
	offset += offset16LEBits

	d.ActiveMode = binary.LittleEndian.Uint32(buf[offset:])
	offset += offset32LEBits

	modes, i, err := readMode(buf, modeCount, offset)
	if err != nil {
		return err
	}

	offset = i
	d.Modes = modes

	zoneCount := binary.LittleEndian.Uint16(buf[offset:])
	offset += offset16LEBits

	zones, i := readZones(buf, zoneCount, offset)
	d.Zones = zones
	offset = i

	ledCount := binary.LittleEndian.Uint16(buf[offset:])
	offset += offset16LEBits

	leds, i, err := readLEDs(buf, ledCount, offset)
	if err != nil {
		return err
	}
	offset = i
	d.LEDs = leds

	colorCount := binary.LittleEndian.Uint16(buf[offset:])
	offset += offset16LEBits

	d.Colors = make([]Color, 0)
	for i := uint16(0); i < colorCount; i++ {
		color, err := readColor(buf, offset)
		if err != nil {
			return err
		}
		d.Colors = append(d.Colors, color)
		offset += 4
	}

	return nil
}

func (d Device) String() string {
	return fmt.Sprintf(`%s (typ %d; ver %s; ser %s)
Mode - Active: %d; Total: %d
	%v
---`,
		d.Name, d.Type, d.Version, d.Serial,
		d.ActiveMode, len(d.Modes), d.Modes[d.ActiveMode])
}
