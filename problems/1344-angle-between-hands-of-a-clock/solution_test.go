package anglebetweenhandsofclock

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func TestAngleClock(t *testing.T) {
	tests := []struct {
		name    string
		hour    int
		minutes int
		want    float64
	}{
		{
			name:    "example 1",
			hour:    12,
			minutes: 30,
			want:    165,
		},
		{
			name:    "example 2",
			hour:    3,
			minutes: 30,
			want:    75,
		},
		{
			name:    "example 3",
			hour:    3,
			minutes: 15,
			want:    7.5,
		},
		{
			name:    "hands overlap at noon",
			hour:    12,
			minutes: 0,
			want:    0,
		},
		{
			name:    "hands point opposite at six",
			hour:    6,
			minutes: 0,
			want:    180,
		},
		{
			name:    "minute hand only movement",
			hour:    12,
			minutes: 15,
			want:    82.5,
		},
		{
			name:    "hour twelve treated as zero",
			hour:    12,
			minutes: 45,
			want:    112.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := angleClock(tt.hour, tt.minutes)
			if !almostEqual(got, tt.want) {
				t.Fatalf("angleClock(%d, %d) = %v, want %v", tt.hour, tt.minutes, got, tt.want)
			}
		})
	}
}
