package countcommasinrange

import "testing"

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "leetcode example 1",
			n:    1002,
			want: 3,
		},
		{
			name: "leetcode example 2",
			n:    998,
			want: 0,
		},
		{
			name: "minimum input",
			n:    1,
			want: 0,
		},
		{
			name: "last value without commas",
			n:    999,
			want: 0,
		},
		{
			name: "first value with one comma",
			n:    1000,
			want: 1,
		},
		{
			name: "upper bound of single-comma range",
			n:    9999,
			want: 9000,
		},
		{
			name: "five digit number",
			n:    10000,
			want: 9001,
		},
		{
			name: "maximum constraint",
			n:    100000,
			want: 99001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCommas(tt.n)
			if got != tt.want {
				t.Fatalf("countCommas(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
