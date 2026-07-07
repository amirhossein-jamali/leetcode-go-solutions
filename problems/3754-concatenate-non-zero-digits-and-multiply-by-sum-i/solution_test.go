package concatenatenonzerodigitsandmultiplybysumi

import "testing"

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{
			name: "leetcode example 1",
			n:    10203004,
			want: 12340,
		},
		{
			name: "leetcode example 2",
			n:    1000,
			want: 1,
		},
		{
			name: "zero input",
			n:    0,
			want: 0,
		},
		{
			name: "single non-zero digit",
			n:    9,
			want: 81,
		},
		{
			name: "no zero digits",
			n:    123,
			want: 738,
		},
		{
			name: "all zeros except leading digit",
			n:    500000,
			want: 25,
		},
		{
			name: "maximum constraint value",
			n:    1_000_000_000,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAndMultiply(tt.n)
			if got != tt.want {
				t.Fatalf("sumAndMultiply(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
