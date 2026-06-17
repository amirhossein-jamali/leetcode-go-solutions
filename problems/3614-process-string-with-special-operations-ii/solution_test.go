package processstringwithspecialoperationsii

import "testing"

func TestProcessStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int64
		want byte
	}{
		{
			name: "example 1",
			s:    "a#b%*",
			k:    1,
			want: 'a',
		},
		{
			name: "example 2",
			s:    "cd%#*#",
			k:    3,
			want: 'd',
		},
		{
			name: "example 3 out of bounds",
			s:    "z*#",
			k:    0,
			want: '.',
		},
		{
			name: "first character",
			s:    "a#b%*",
			k:    0,
			want: 'b',
		},
		{
			name: "k beyond final length",
			s:    "a#b%*",
			k:    2,
			want: '.',
		},
		{
			name: "single letter at index zero",
			s:    "x",
			k:    0,
			want: 'x',
		},
		{
			name: "empty result any k",
			s:    "z*",
			k:    0,
			want: '.',
		},
		{
			name: "duplicate maps second half index to first half",
			s:    "ab#",
			k:    2,
			want: 'a',
		},
		{
			name: "reverse flips index",
			s:    "ab%",
			k:    0,
			want: 'b',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processStr(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("processStr(%q, %d) = %q, want %q", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
