package openrgb

type LED struct {
	Name  string
	Value Color
}

func readLEDs(buf []byte, count uint16, offset int) ([]LED, int, error) {
	leds := make([]LED, 0)

	for ledIndex := uint16(0); ledIndex < count; ledIndex++ {
		name, i := readString(buf, offset)
		offset += i
		color, err := readColor(buf, offset)
		if err != nil {
			return nil, 0, err
		}

		offset += 4

		leds = append(leds, LED{
			Name:  name,
			Value: color,
		})
	}

	return leds, offset, nil
}
