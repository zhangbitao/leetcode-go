package algorithms

import (
	"testing"
)

func buildList(value []int) *ListNode {
	var result *ListNode
	r := &result
	for _, v := range value {
		l := &ListNode{
			Val:  v,
			Next: nil,
		}
		*r = l
		r = &(l.Next)
	}
	return result
}

func TestNumComponents(t *testing.T) {
	testCases := []struct {
		input  []int
		G      []int
		output int
	}{
		{[]int{0, 1, 2, 3}, []int{0, 1, 3}, 2},
		{[]int{0, 1, 2, 3, 4}, []int{0, 3, 1, 4}, 2},
		{[]int{0}, []int{0}, 1},
		{[]int{0, 1, 2}, []int{1, 0}, 1},
	}

	for _, tc := range testCases {
		head := buildList(tc.input)
		got := numComponents(head, tc.G)
		want := tc.output
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
