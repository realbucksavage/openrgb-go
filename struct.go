package openrgb

type orgbHeader struct {
	deviceID  uint32
	commandID uint32
	length    uint32
}

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
	Zone        []Zone
}

type LED struct {
	Name  string
	Value Color
}

type Color struct {
	Red   int
	Green int
	Blue  int
}

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

type Zone struct {
	Name      string
	Type      int
	MinLEDs   int
	MaxLEDs   int
	TotalLEDs int
}
