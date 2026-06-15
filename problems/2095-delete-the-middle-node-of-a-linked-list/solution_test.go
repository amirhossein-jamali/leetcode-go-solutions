package deletemiddlenode

import (
	"reflect"
	"testing"
)

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

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{
			name: "example 1",
			vals: []int{1, 3, 4, 7, 1, 2, 6},
			want: []int{1, 3, 4, 1, 2, 6},
		},
		{
			name: "example 2",
			vals: []int{1, 2, 3, 4},
			want: []int{1, 2, 4},
		},
		{
			name: "example 3",
			vals: []int{2, 1},
			want: []int{2},
		},
		{
			name: "single node returns nil",
			vals: []int{42},
			want: nil,
		},
		{
			name: "three nodes deletes index 1",
			vals: []int{1, 2, 3},
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := intSliceToList(tt.vals)
			got := deleteMiddle(head)
			gotVals := listToSlice(got)
			if !reflect.DeepEqual(gotVals, tt.want) {
				t.Fatalf("deleteMiddle() = %v, want %v", gotVals, tt.want)
			}
		})
	}
}
