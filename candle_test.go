package candlePrintGo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOHCLArray(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		wantOpen  float64
		wantHigh  float64
		wantLow   float64
		wantClose float64
		wantBull  bool
		wantTop   float64
		wantBot   float64
	}{
		{
			name:      "bullish candle (close > open)",
			input:     []float64{100, 110, 90, 105},
			wantOpen:  100,
			wantHigh:  110,
			wantLow:   90,
			wantClose: 105,
			wantBull:  true,
			wantTop:   105, // close
			wantBot:   100, // open
		},
		{
			name:      "bearish candle (open > close)",
			input:     []float64{105, 110, 90, 100},
			wantOpen:  105,
			wantHigh:  110,
			wantLow:   90,
			wantClose: 100,
			wantBull:  false,
			wantTop:   105, // open
			wantBot:   100, // close
		},
		{
			name:      "doji candle (open == close)",
			input:     []float64{100, 110, 90, 100},
			wantOpen:  100,
			wantHigh:  110,
			wantLow:   90,
			wantClose: 100,
			wantBull:  false,
			wantTop:   100, // open (not bullish)
			wantBot:   100, // close
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ohcl := NewOHCLArrayFromArray(tt.input)

			require.Equal(t, tt.wantOpen, ohcl.Open(), "Open()")
			require.Equal(t, tt.wantHigh, ohcl.High(), "High()")
			require.Equal(t, tt.wantLow, ohcl.Low(), "Low()")
			require.Equal(t, tt.wantClose, ohcl.Close(), "Close()")
			require.Equal(t, tt.wantBull, ohcl.IsBullish(), "IsBullish()")
			require.Equal(t, tt.wantTop, ohcl.Top(), "Top()")
			require.Equal(t, tt.wantBot, ohcl.Bottom(), "Bottom()")
		})
	}
}
