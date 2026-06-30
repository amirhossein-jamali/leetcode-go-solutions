package numberofsubstringscontainingallthreecharacters

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabc",
			want: 10,
		},
		{
			name: "example 2",
			s:    "aaacb",
			want: 3,
		},
		{
			name: "example 3",
			s:    "abc",
			want: 1,
		},
		{
			name: "missing one character",
			s:    "aaaa",
			want: 0,
		},
		{
			name: "repeated pattern",
			s:    "abcabcabc",
			want: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfSubstrings(tt.s)
			if got != tt.want {
				t.Fatalf("numberOfSubstrings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
