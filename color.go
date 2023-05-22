package candlePrintGo

import "github.com/muesli/termenv"

const (
	_bearColor = "#E88388"
	_bullColor = "#A8CC8C"
)

type ColorProfile struct {
	p         termenv.Profile
	bearColor string
	bullColor string
}

var DefaultColorScheme = ColorProfile{
	p:         termenv.ColorProfile(),
	bearColor: _bearColor,
	bullColor: _bullColor,
}
