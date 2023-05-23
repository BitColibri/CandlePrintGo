package candlePrintGo

import (
	"fmt"
	"math"

	"github.com/muesli/termenv"
)

const (
	_symbolStick            = "│"
	_symbolCandle           = "┃"
	_symbolHalfTop          = "╽"
	_symbolHalfBottom       = "╿"
	_symbolHalfCandleTop    = "╻"
	_symbolHalfCandleBottom = "╹"
	_symbolHalfStickTop     = "╷"
	_symbolHalfStickBottom  = "╵"
	_symbolEmpty            = " "

	_halfTop    = 0.75
	_halfBottom = 0.25
)

type Chart interface {
	Render() string
}

var _ Chart = &CandleChart{}

type CandleChart struct {
	profile     *ColorProfile
	height      float64
	data        []Candle
	chartBottom float64
	chartTop    float64
}

type Option func(chart *CandleChart)

func WithColorProfile(profile *ColorProfile) Option {
	return func(c *CandleChart) {
		c.profile = profile
	}
}

func NewCandleChart(data []Candle, height float64, opts ...Option) *CandleChart {
	min := math.MaxFloat64
	max := math.SmallestNonzeroFloat64

	for _, v := range data {
		min = math.Min(min, v.Bottom())
		max = math.Max(max, v.Top())
	}

	chart := &CandleChart{
		profile:     DefaultColorScheme,
		height:      height,
		data:        data,
		chartBottom: min,
		chartTop:    max,
	}
	for _, o := range opts {
		o(chart)
	}

	return chart
}

func (c *CandleChart) toHeightUnits(x float64) float64 {
	return (x - c.chartBottom) / (c.chartTop - c.chartBottom) * c.height
}

func (c *CandleChart) renderCandle(tick Candle, height float64) string {
	top := c.toHeightUnits(tick.High())
	topCandle := c.toHeightUnits(tick.Top())

	bottom := c.toHeightUnits(tick.Low())
	bottomCandle := c.toHeightUnits(tick.Bottom())

	if math.Ceil(top) > height && height >= math.Floor(topCandle) {
		if topCandle-height > _halfTop {
			return c.colorCandle(_symbolCandle, tick.IsBullish())
		} else if topCandle-height > _halfBottom {
			if top-height > _halfTop {
				return c.colorCandle(_symbolHalfTop, tick.IsBullish())
			}
			return c.colorCandle(_symbolHalfCandleTop, tick.IsBullish())
		} else {
			if top-height > _halfTop {
				return c.colorCandle(_symbolStick, tick.IsBullish())
			} else if top-height > _halfBottom {
				return c.colorCandle(_symbolHalfStickTop, tick.IsBullish())
			}
			return _symbolEmpty
		}
	} else if math.Floor(topCandle) >= height && height >= math.Ceil(bottomCandle) {
		return c.colorCandle(_symbolCandle, tick.IsBullish())
	} else if math.Ceil(bottomCandle) >= height && height >= math.Floor(bottom) {
		if bottomCandle-height < _halfBottom {
			return c.colorCandle(_symbolCandle, tick.IsBullish())
		} else if bottomCandle-height < _halfTop {
			if bottom-height < _halfBottom {
				return c.colorCandle(_symbolHalfBottom, tick.IsBullish())
			}
			return c.colorCandle(_symbolHalfCandleBottom, tick.IsBullish())
		} else {
			if bottom-height < _halfBottom {
				return c.colorCandle(_symbolStick, tick.IsBullish())
			} else if bottom-height < _halfTop {
				return c.colorCandle(_symbolHalfStickBottom, tick.IsBullish())
			}
			return _symbolEmpty
		}
	}
	return _symbolEmpty
}

func (c *CandleChart) colorCandle(symbol string, isBulish bool) string {
	s := termenv.String(symbol).Foreground(c.profile.p.Color(c.profile.bearColor))
	if isBulish {
		s = termenv.String(symbol).Foreground(c.profile.p.Color(c.profile.bullColor))
	}
	return fmt.Sprintf(`%v`, s)
}

func (c *CandleChart) Render() string {
	r := "\n"
	for i := c.height; i >= 0; i-- {
		for _, v := range c.data {
			r += c.renderCandle(v, float64(i))
		}
		r += "\n"
	}
	return r
}
