package candlePrintGo

type Candle interface {
	Open() float64
	High() float64
	Low() float64
	Close() float64
	// Top returns what is the top of the candle either Open or Close
	Top() float64
	// Bottom returns what is the bottom of the candle either Open or Close
	Bottom() float64
	// IsBullish check the tendency of the candle
	IsBullish() bool
}

type CandleBar struct {
	isBullish bool
	open      float64
	high      float64
	low       float64
	close     float64
}

func NewCandleBar(o, h, l, c float64) *CandleBar {
	return &CandleBar{
		isBullish: c > o,
		open:      o,
		high:      h,
		low:       l,
		close:     c,
	}
}

func (c CandleBar) Open() float64 {
	return c.close
}

func (c CandleBar) High() float64 {
	return c.high
}

func (c CandleBar) Low() float64 {
	return c.low
}

func (c CandleBar) Close() float64 {
	return c.close
}

func (c CandleBar) IsBullish() bool {
	return c.isBullish
}

func (c CandleBar) Top() float64 {
	if c.IsBullish() {
		return c.close
	}
	return c.open
}

func (c CandleBar) Bottom() float64 {
	if c.IsBullish() {
		return c.open
	}
	return c.close
}
