package bubbleCandle

import (
	"fmt"
	"math"

	"github.com/muesli/termenv"
)

const (
	SYMBOL_STICK              = "│"
	SYMBOL_CANDLE             = "┃"
	SYMBOL_HALF_TOP           = "╽"
	SYMBOL_HALF_BOTTOM        = "╿"
	SYMBOL_HALF_CANDLE_TOP    = "╻"
	SYMBOL_HALF_CANDLE_BOTTOM = "╹"
	SYMBOL_HALF_STICK_TOP     = "╷"
	SYMBOL_HALF_STICK_BOTTOM  = "╵"
	SYMBOL_NOTHING            = " "
)

type Candle struct {
	IsBullish bool
	Open      float64
	High      float64
	Low       float64
	Close     float64
}

func NewCandle(o, h, l, c float64) *Candle {
	return &Candle{
		IsBullish: c > o,
		Open:      o,
		High:      h,
		Low:       l,
		Close:     c,
	}
}

func (c Candle) Top() float64 {
	if c.IsBullish {
		return c.Close
	}
	return c.Open
}

func (c Candle) Bottom() float64 {
	if c.IsBullish {
		return c.Open
	}
	return c.Close
}

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
			return c.colorCandle(SYMBOL_CANDLE, tick.IsBullish)
		} else if topCandle-height > 0.25 {
			if top-height > 0.75 {
				return c.colorCandle(SYMBOL_HALF_TOP, tick.IsBullish)
			}
			return c.colorCandle(SYMBOL_HALF_CANDLE_TOP, tick.IsBullish)
		} else {
			if top-height > 0.75 {
				return c.colorCandle(SYMBOL_STICK, tick.IsBullish)
			} else if top-height > 0.25 {
				return c.colorCandle(SYMBOL_HALF_STICK_TOP, tick.IsBullish)
			}
			return SYMBOL_NOTHING
		}
	} else if math.Floor(topCandle) >= height && height >= math.Ceil(bottomCandle) {
		return c.colorCandle(SYMBOL_CANDLE, tick.IsBullish)
	} else if math.Ceil(bottomCandle) >= height && height >= math.Floor(bottom) {
		if bottomCandle-height < 0.25 {
			return c.colorCandle(SYMBOL_CANDLE, tick.IsBullish)
		} else if bottomCandle-height < 0.75 {
			if bottom-height < 0.25 {
				return c.colorCandle(SYMBOL_HALF_BOTTOM, tick.IsBullish)
			}
			return c.colorCandle(SYMBOL_HALF_CANDLE_BOTTOM, tick.IsBullish)
		} else {
			if bottom-height < 0.25 {
				return c.colorCandle(SYMBOL_STICK, tick.IsBullish)
			} else if bottom-height < 0.75 {
				return c.colorCandle(SYMBOL_HALF_STICK_BOTTOM, tick.IsBullish)
			}
			return SYMBOL_NOTHING
		}
	}
	return SYMBOL_NOTHING
}

func (c *CandleChart) colorCandle(symbol string, isBulish bool) string {
	s := termenv.String(symbol).Foreground(c.profile.Color("#E88388"))
	if isBulish {
		s = termenv.String(symbol).Foreground(c.profile.Color("#A8CC8C"))
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
