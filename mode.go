package openrgb

import (
	"encoding/binary"
	"fmt"
)

type Mode struct {
	Name      string
	Value     uint32
	Flags     uint32
	MinSpeed  uint32
	MaxSpeed  uint32
	MinColors uint32
	MaxColors uint32
	Speed     uint32
	Direction uint32
	ColorMode uint32
	Colors    []Color
}

func readMode(buf []byte, modeCount uint16, offset int) ([]Mode, int, error) {
	modes := make([]Mode, 0)
	colors := make([]Color, 0)

	for modeIndex := uint16(0); modeIndex < modeCount; modeIndex++ {
		modeName, i := readString(buf, offset)
		offset += i

		mode := Mode{Name: modeName}
		for _, ptr := range []*uint32{
			&mode.Value,
			&mode.Flags,
			&mode.MinSpeed,
			&mode.MaxSpeed,
			&mode.MinColors,
			&mode.MaxColors,
			&mode.Speed,
			&mode.Direction,
			&mode.ColorMode,
		} {
			*ptr = binary.LittleEndian.Uint32(buf[offset:])
			offset += 4
		}

		colorLength := binary.LittleEndian.Uint16(buf[offset:])
		offset += 2

		var ci uint16 = 0
		for ; ci < colorLength; ci++ {
			color, err := readColor(buf, offset)
			if err != nil {
				return nil, 0, err
			}
			offset += 4
			colors = append(colors, color)
		}

		mode.Colors = colors

		modes = append(modes, mode)
	}

	return modes, offset, nil
}
func (m Mode) String() string {
	return fmt.Sprintf(`%s
	Speed : %d (%d - %d)
	ColorMode : %s
	Colors: %v`,
		m.Name,
		m.Speed, m.MinSpeed, m.MaxSpeed,
		colorMode(m.ColorMode),
		m.Colors)
}

func colorMode(mode uint32) string {
	switch mode {
	case 1:
		return "Per-LED"
	case 2:
		return "Mode-Specific"
	case 3:
		return "Random"
	default:
		return "Unidentified"
	}
}
