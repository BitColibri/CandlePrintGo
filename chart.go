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

	_bearColor = "#E88388"
	_bullColor = "#A8CC8C"
)

type CandleChart struct {
	profile     termenv.Profile
	height      float64
	data        []*Candle
	chartBottom float64
	chartTop    float64
}

func NewCandleChart(data []*Candle, height float64, p termenv.Profile) *CandleChart {
	min := math.MaxFloat64
	max := math.SmallestNonzeroFloat64

	for _, v := range data {
		min = math.Min(min, v.Bottom())
		max = math.Max(max, v.Top())
	}

	return &CandleChart{
		profile:     p,
		height:      height,
		data:        data,
		chartBottom: min,
		chartTop:    max,
	}
}

func (c *CandleChart) Data(data []*Candle) {
	c.data = data
}

func (c *CandleChart) toHeightUnits(x float64) float64 {
	return (x - c.chartBottom) / (c.chartTop - c.chartBottom) * c.height
}

func (c *CandleChart) renderCandle(tick *Candle, height float64) string {
	top := c.toHeightUnits(tick.High)
	topCandle := c.toHeightUnits(tick.Top())

	bottom := c.toHeightUnits(tick.Low)
	bottomCandle := c.toHeightUnits(tick.Bottom())

	if math.Ceil(top) > height && height >= math.Floor(topCandle) {
		if topCandle-height > 0.75 {
			return c.colorCandle(_symbolCandle, tick.IsBullish)
		} else if topCandle-height > 0.25 {
			if top-height > 0.75 {
				return c.colorCandle(_symbolHalfTop, tick.IsBullish)
			}
			return c.colorCandle(_symbolHalfCandleTop, tick.IsBullish)
		} else {
			if top-height > 0.75 {
				return c.colorCandle(_symbolStick, tick.IsBullish)
			} else if top-height > 0.25 {
				return c.colorCandle(_symbolHalfStickTop, tick.IsBullish)
			}
			return _symbolEmpty
		}
	} else if math.Floor(topCandle) >= height && height >= math.Ceil(bottomCandle) {
		return c.colorCandle(_symbolCandle, tick.IsBullish)
	} else if math.Ceil(bottomCandle) >= height && height >= math.Floor(bottom) {
		if bottomCandle-height < 0.25 {
			return c.colorCandle(_symbolCandle, tick.IsBullish)
		} else if bottomCandle-height < 0.75 {
			if bottom-height < 0.25 {
				return c.colorCandle(_symbolHalfBottom, tick.IsBullish)
			}
			return c.colorCandle(_symbolHalfCandleBottom, tick.IsBullish)
		} else {
			if bottom-height < 0.25 {
				return c.colorCandle(_symbolStick, tick.IsBullish)
			} else if bottom-height < 0.75 {
				return c.colorCandle(_symbolHalfStickBottom, tick.IsBullish)
			}
			return _symbolEmpty
		}
	}
	return _symbolEmpty
}

func (c *CandleChart) colorCandle(symbol string, isBulish bool) string {
	s := termenv.String(symbol).Foreground(c.profile.Color(_bearColor))
	if isBulish {
		s = termenv.String(symbol).Foreground(c.profile.Color(_bullColor))
	}
	return fmt.Sprintf(`%v`, s)
}

func (c *CandleChart) Render() string {
	r := "\n"
	for i := c.height; i >= 0; i-- {
		//if i%4 == 0 {
		//	//calc := c.chartBottom + (float64(i) * (c.chartTop - c.chartBottom) / c.height)
		//	//
		//	//r += fmt.Sprintf("%.5f", calc)
		//	r += "         "
		//} else {
		//	r += "         "
		//}
		for _, v := range c.data {
			r += c.renderCandle(v, float64(i))
		}
		r += "\n"
	}
	return r
}
