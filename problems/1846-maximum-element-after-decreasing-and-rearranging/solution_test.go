package maximumelementafterdecreasingandrearranging

import "testing"

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "example 1",
			arr:  []int{2, 2, 1, 2, 1},
			want: 2,
		},
		{
			name: "example 2",
			arr:  []int{100, 1, 1000},
			want: 3,
		},
		{
			name: "example 3 already valid",
			arr:  []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "single element",
			arr:  []int{9},
			want: 1,
		},
		{
			name: "many ones",
			arr:  []int{1, 1, 1, 1},
			want: 1,
		},
		{
			name: "two elements",
			arr:  []int{5, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.arr...)
			got := maximumElementAfterDecrementingAndRearranging(arr)
			if got != tt.want {
				t.Fatalf(
					"maximumElementAfterDecrementingAndRearranging(%v) = %d, want %d",
					tt.arr,
					got,
					tt.want,
				)
			}
		})
	}
}
