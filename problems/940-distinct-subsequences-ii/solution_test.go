package distinctsubsequencesii

import "testing"

func TestDistinctSubseqII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "leetcode example 1",
			s:    "abc",
			want: 7,
		},
		{
			name: "leetcode example 2",
			s:    "aba",
			want: 6,
		},
		{
			name: "leetcode example 3",
			s:    "aaa",
			want: 3,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "two distinct characters",
			s:    "ab",
			want: 3,
		},
		{
			name: "repeated pair",
			s:    "aab",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distinctSubseqII(tt.s)
			if got != tt.want {
				t.Fatalf("distinctSubseqII(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
