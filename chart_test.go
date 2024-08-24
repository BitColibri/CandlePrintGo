package candlePrintGo

import "testing"

func Test_formatPrice(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{
			name:  "check one",
			value: 1,
			want:  "    1.00 ",
		},
		{
			name:  "check 124.23",
			value: 124.23,
			want:  "  124.23 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPrice(tt.value); got != tt.want {
				t.Errorf("formatFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
