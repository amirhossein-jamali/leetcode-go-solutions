package rearrangecharacterstomaketargetstring

import "testing"

func TestRearrangeCharacters(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		target string
		want   int
	}{
		{
			name:   "example 1",
			s:      "ilovecodingonleetcode",
			target: "code",
			want:   2,
		},
		{
			name:   "example 2",
			s:      "abcba",
			target: "abc",
			want:   1,
		},
		{
			name:   "example 3",
			s:      "abbaccaddaeea",
			target: "aaaaa",
			want:   1,
		},
		{
			name:   "balloon equivalent of 1189 example 1",
			s:      "nlaebolko",
			target: "balloon",
			want:   1,
		},
		{
			name:   "balloon equivalent of 1189 example 2",
			s:      "loonbalxballpoon",
			target: "balloon",
			want:   2,
		},
		{
			name:   "no copies possible",
			s:      "leetcode",
			target: "balloon",
			want:   0,
		},
		{
			name:   "exact single match",
			s:      "code",
			target: "code",
			want:   1,
		},
		{
			name:   "single character target",
			s:      "aaaa",
			target: "a",
			want:   4,
		},
		{
			name:   "missing required letter",
			s:      "abc",
			target: "abcd",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rearrangeCharacters(tt.s, tt.target)
			if got != tt.want {
				t.Fatalf("rearrangeCharacters(%q, %q) = %d, want %d", tt.s, tt.target, got, tt.want)
			}
		})
	}
}
