package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "example 3 spaces only",
			s:    " ",
			want: true,
		},
		{
			name: "alphanumeric only true",
			s:    "aba",
			want: true,
		},
		{
			name: "alphanumeric only false",
			s:    "ab",
			want: false,
		},
		{
			name: "digits and letters",
			s:    "0P",
			want: false,
		},
		{
			name: "empty after strip is palindrome via single pass",
			s:    ".,",
			want: true,
		},
		{
			name: "mixed case same letters",
			s:    "Aa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Fatalf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
