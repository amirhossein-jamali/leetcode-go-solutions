package processstringwithspecialoperationsi

import "testing"

func TestProcessStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example 1",
			s:    "a#b%*",
			want: "ba",
		},
		{
			name: "example 2",
			s:    "z*#",
			want: "",
		},
		{
			name: "single letter",
			s:    "a",
			want: "a",
		},
		{
			name: "duplicate only",
			s:    "a#",
			want: "aa",
		},
		{
			name: "reverse only",
			s:    "ab%",
			want: "ba",
		},
		{
			name: "delete on empty result is no-op",
			s:    "*",
			want: "",
		},
		{
			name: "duplicate empty result stays empty",
			s:    "#",
			want: "",
		},
		{
			name: "reverse empty result stays empty",
			s:    "%",
			want: "",
		},
		{
			name: "append then duplicate",
			s:    "ab#",
			want: "abab",
		},
		{
			name: "duplicate then reverse",
			s:    "a#%",
			want: "aa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processStr(tt.s)
			if got != tt.want {
				t.Fatalf("processStr(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
