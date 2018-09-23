package algorithms

import (
	"testing"
)

func arrayToLists(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &ListNode{
		Val: 0,
	}
	prev := head
	for _, a := range array {
		l := &ListNode{
			Val: a,
		}
		prev.Next = l
		prev = prev.Next
	}

	return head.Next
}

func isSameLists(l1 *ListNode, l2 *ListNode) bool {
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

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
	}

	for _, tc := range testCases {
		head := arrayToLists(tc.input)
		got := swapPairs(head)
		want := arrayToLists(tc.output)
		if !isSameLists(got, want) {
			t.Errorf("wrong result, expected %v and got %v", printList(want), printList(got))
		}
	}
}
