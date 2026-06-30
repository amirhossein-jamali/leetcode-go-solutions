package numberofstringsthatappearsassubstringsinword

import "testing"

func TestNumOfStrings(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		word     string
		want     int
	}{
		{
			name:     "example 1",
			patterns: []string{"a", "abc", "bc", "d"},
			word:     "abc",
			want:     3,
		},
		{
			name:     "example 2",
			patterns: []string{"a", "b", "c"},
			word:     "aaaaabbbbb",
			want:     2,
		},
		{
			name:     "example 3 duplicate patterns",
			patterns: []string{"a", "a", "a"},
			word:     "ab",
			want:     3,
		},
		{
			name:     "no matches",
			patterns: []string{"x", "y", "z"},
			word:     "abc",
			want:     0,
		},
		{
			name:     "exact word match",
			patterns: []string{"hello"},
			word:     "hello",
			want:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfStrings(tt.patterns, tt.word)
			if got != tt.want {
				t.Fatalf(
					"numOfStrings(%v, %q) = %d, want %d",
					tt.patterns,
					tt.word,
					got,
					tt.want,
				)
			}
		})
	}
}
