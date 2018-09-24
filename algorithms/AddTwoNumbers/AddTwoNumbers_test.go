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

func printList(l *ListNode) []int {
	var r []int
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		v1 []int
		v2 []int
		r  []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{6, 2, 7, 4}, []int{2, 1, 9}, []int{8, 3, 6, 5}},
		{[]int{9, 8}, []int{3, 9, 9, 1, 9}, []int{2, 8, 0, 2, 9}},
	}

	for _, tc := range testCases {
		l1 := buildList(tc.v1)
		l2 := buildList(tc.v2)
		got := addTwoNumbers(l1, l2)
		want := buildList(tc.r)
		if !isSameList(got, want) {
			t.Errorf("wrong result, expected %v and got %v", printList(want), printList(got))
		}
	}
}
