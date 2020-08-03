package openrgb

type Device struct {
	Type        int
	Name        string
	Description string
	Version     string
	Serial      string
	Location    string
	ActiveMode  int
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
	Value     int
	Flags     int
	MinSpeed  int
	MaxSpeed  int
	MinColors int
	MaxColors int
	Speed     int
	Direction int
	ColorMode int
	Colors    []Color
}

type Zone struct {
	Name      string
	Type      int
	MinLEDs   int
	MaxLEDs   int
	TotalLEDs int
}
