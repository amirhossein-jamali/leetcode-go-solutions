package distinctsubsequences

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "leetcode example 1",
			s:    "rabbbit",
			t:    "rabbit",
			want: 3,
		},
		{
			name: "leetcode example 2",
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
		{
			name: "exact match",
			s:    "abc",
			t:    "abc",
			want: 1,
		},
		{
			name: "target longer than source",
			s:    "ab",
			t:    "abc",
			want: 0,
		},
		{
			name: "no matching subsequence",
			s:    "abc",
			t:    "def",
			want: 0,
		},
		{
			name: "single character repeated in source",
			s:    "aaa",
			t:    "a",
			want: 3,
		},
		{
			name: "multiple ways to skip characters",
			s:    "aab",
			t:    "ab",
			want: 2,
		},
		{
			name: "single character source and target",
			s:    "a",
			t:    "a",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDistinct(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("numDistinct(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
