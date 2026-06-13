package weightedwordmapping

import "testing"

func TestMapWordWeights(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		weights []int
		want    string
	}{
		{
			name:  "example 1",
			words: []string{"abcd", "def", "xyz"},
			weights: []int{
				5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7,
				8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2,
			},
			want: "rij",
		},
		{
			name:  "example 2",
			words: []string{"a", "b", "c"},
			weights: []int{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			},
			want: "yyy",
		},
		{
			name:  "example 3",
			words: []string{"abcd"},
			weights: []int{
				7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10,
				2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5,
			},
			want: "g",
		},
		{
			name:  "modulo zero maps to z",
			words: []string{"az"},
			weights: []int{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 25,
			},
			want: "z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapWordWeights(tt.words, tt.weights)
			if got != tt.want {
				t.Fatalf("mapWordWeights() = %q, want %q", got, tt.want)
			}
		})
	}
}
