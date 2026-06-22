package maximumnumberofballoons

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			name: "example 1",
			text: "nlaebolko",
			want: 1,
		},
		{
			name: "example 2",
			text: "loonbalxballpoon",
			want: 2,
		},
		{
			name: "example 3",
			text: "leetcode",
			want: 0,
		},
		{
			name: "exactly one balloon",
			text: "balloon",
			want: 1,
		},
		{
			name: "two complete balloons",
			text: "balloonballoon",
			want: 2,
		},
		{
			name: "missing b",
			text: "alloon",
			want: 0,
		},
		{
			name: "limited by double l",
			text: "balon",
			want: 0,
		},
		{
			name: "single character",
			text: "b",
			want: 0,
		},
		{
			name: "extra unrelated letters",
			text: "balloonzzzz",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxNumberOfBalloons(tt.text)
			if got != tt.want {
				t.Fatalf("maxNumberOfBalloons(%q) = %d, want %d", tt.text, got, tt.want)
			}
		})
	}
}
