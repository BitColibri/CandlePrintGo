package candlePrintGo

// Candle represents a financial candlestick with OHLC (Open, High, Low, Close) data.
type Candle interface {
	Open() float64
	High() float64
	Low() float64
	Close() float64
	// Top returns the top of the candle body (either Open or Close, whichever is higher).
	Top() float64
	// Bottom returns the bottom of the candle body (either Open or Close, whichever is lower).
	Bottom() float64
	// IsBullish returns true if the candle is bullish (Close > Open).
	IsBullish() bool
}

var _ Candle = &CandleBar{}
var _ Candle = &OHCLArray{}

// OHCLArray is a fixed-size array representing a candle with [Open, High, Low, Close] values.
type OHCLArray [4]float64

func (o OHCLArray) Open() float64 {
	return o[0]
}

func (o OHCLArray) High() float64 {
	return o[1]
}

func (o OHCLArray) Low() float64 {
	return o[2]
}

func (o OHCLArray) Close() float64 {
	return o[3]
}

func (o OHCLArray) IsBullish() bool {
	return o[3] > o[0]
}

func (o OHCLArray) Top() float64 {
	if o.IsBullish() {
		return o[3] // close
	}
	return o[0] // open
}

func (o OHCLArray) Bottom() float64 {
	if o.IsBullish() {
		return o[0] // open
	}
	return o[3] // close
}

// NewOHCLArray creates a new OHCLArray from individual Open, High, Low, Close values.
func NewOHCLArray(o, h, l, c float64) *OHCLArray {
	return &OHCLArray{o, h, l, c}
}

// NewOHCLArrayFromArray creates a new OHCLArray from a slice where elements are [Open, High, Low, Close].
func NewOHCLArrayFromArray(arr []float64) *OHCLArray {
	return &OHCLArray{arr[0], arr[1], arr[2], arr[3]}
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

func NewCandleBarFromArray(arr []float64) *CandleBar {
	return &CandleBar{
		isBullish: arr[3] > arr[0],
		open:      arr[0],
		high:      arr[1],
		low:       arr[2],
		close:     arr[3],
	}
}
func (c CandleBar) Open() float64 {
	return c.open
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
