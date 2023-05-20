package candlePrintGo

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
