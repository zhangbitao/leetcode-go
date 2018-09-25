package algorithms

import (
	"testing"
)

func buildCycle(value []int, isCycle bool) *ListNode {
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
	if isCycle {
		*r = result
	}
	return result
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		input   []int
		isCycle bool
		output  bool
	}{
		{[]int{1, 1, 2}, true, true},
		{[]int{1, 2, 3}, false, false},
	}

	for _, tc := range testCases {
		head := buildCycle(tc.input, tc.isCycle)
		got := hasCycle(head)
		want := tc.output
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
