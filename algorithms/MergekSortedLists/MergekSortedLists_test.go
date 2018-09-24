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

func isSameList(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	for l1 != nil {
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		return false
	}

	return true
}

func printList(l *ListNode) (r []int) {
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return
}

func TestMergeKLists(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output []int
	}{
		{[][]int{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{[]int{1, 2, 3}, []int{}, []int{1, 3, 5}}, []int{1, 1, 2, 3, 3, 5}},
	}

	for _, tc := range testCases {
		var lists []*ListNode
		for _, l := range tc.input {
			lists = append(lists, buildList(l))
		}
		got := mergeKLists(lists)
		want := buildList(tc.output)
		if !isSameList(got, want) {
			t.Errorf("wrong result, expected %v and got %v", printList(want), printList(got))
		}
	}
}
