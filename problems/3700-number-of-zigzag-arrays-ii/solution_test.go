package numberofzigzagarraysii

import "testing"

func TestZigZagArrays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		l    int
		r    int
		want int
	}{
		{
			name: "example 1",
			n:    3,
			l:    4,
			r:    5,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			l:    1,
			r:    3,
			want: 10,
		},
		{
			name: "two value range any length",
			n:    100,
			l:    1,
			r:    2,
			want: 2,
		},
		{
			name: "length five matches part one",
			n:    5,
			l:    1,
			r:    3,
			want: 26,
		},
		{
			name: "even length four",
			n:    4,
			l:    1,
			r:    3,
			want: 16,
		},
		{
			name: "large n small range",
			n:    1_000_000_000,
			l:    1,
			r:    3,
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zigZagArrays(tt.n, tt.l, tt.r)
			if got != tt.want {
				t.Fatalf("zigZagArrays(%d, %d, %d) = %d, want %d", tt.n, tt.l, tt.r, got, tt.want)
			}
		})
	}
}
