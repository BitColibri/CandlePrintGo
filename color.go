package candlePrintGo

import "github.com/muesli/termenv"

const (
	_bearColor = "#E88388"
	_bullColor = "#A8CC8C"
)

type ColorProfile struct {
	p         termenv.Profile
	bullColor string
	bearColor string
}

func NewColorProfile(profile termenv.Profile, bullColor, bearColor string) *ColorProfile {
	return &ColorProfile{
		p:         profile,
		bullColor: bullColor,
		bearColor: bearColor,
	}
}

var DefaultColorScheme = &ColorProfile{
	p:         termenv.ColorProfile(),
	bullColor: _bullColor,
	bearColor: _bearColor,
}
