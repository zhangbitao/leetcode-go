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

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{1, 2, 4}, []int{}, []int{1, 2, 4}},
		{[]int{}, []int{1, 3, 4}, []int{1, 3, 4}},
		{[]int{1, 1, 1}, []int{1, 2, 3, 4, 5}, []int{1, 1, 1, 1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		l1 := arrayToLists(tc.input1)
		l2 := arrayToLists(tc.input2)
		got := mergeTwoLists(l1, l2)
		want := arrayToLists(tc.output)
		if !isSameLists(got, want) {
			t.Errorf("wrong result, expected %v and got %v", printList(want), printList(got))
		}
	}
}
