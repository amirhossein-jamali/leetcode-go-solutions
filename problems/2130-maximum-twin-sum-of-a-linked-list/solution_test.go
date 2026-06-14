package maximumtwinsum

import "testing"

func intSliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func TestPairSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{
			name: "example 1",
			vals: []int{5, 4, 2, 1},
			want: 6,
		},
		{
			name: "example 2",
			vals: []int{4, 2, 2, 3},
			want: 7,
		},
		{
			name: "example 3",
			vals: []int{1, 100000},
			want: 100001,
		},
		{
			name: "minimum length two nodes",
			vals: []int{1, 2},
			want: 3,
		},
		{
			name: "six nodes alternating peaks",
			vals: []int{1, 2, 3, 4, 5, 6},
			want: 7,
		},
		{
			name: "all equal twin sums",
			vals: []int{3, 3, 3, 3},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := intSliceToList(tt.vals)
			got := pairSum(head)
			if got != tt.want {
				t.Fatalf("pairSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
