package sequentialdigits

import (
	"reflect"
	"testing"
)

func TestSequentialDigits(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want []int
	}{
		{
			name: "leetcode example 1",
			low:  100,
			high: 300,
			want: []int{123, 234},
		},
		{
			name: "leetcode example 2",
			low:  1000,
			high: 13000,
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
		{
			name: "two digit range",
			low:  12,
			high: 89,
			want: []int{12, 23, 34, 45, 56, 67, 78, 89},
		},
		{
			name: "single matching value",
			low:  123,
			high: 123,
			want: []int{123},
		},
		{
			name: "no sequential digits in range",
			low:  10,
			high: 11,
			want: []int{},
		},
		{
			name: "range ends before first match",
			low:  100,
			high: 122,
			want: []int{},
		},
		{
			name: "includes longest sequential number",
			low:  123456789,
			high: 123456789,
			want: []int{123456789},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequentialDigits(tt.low, tt.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"sequentialDigits(%d, %d) = %v, want %v",
					tt.low,
					tt.high,
					got,
					tt.want,
				)
			}
		})
	}
}
