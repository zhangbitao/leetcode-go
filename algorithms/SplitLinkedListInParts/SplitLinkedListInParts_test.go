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

func isSameList(l1 []*ListNode, l2 []*ListNode) bool {
	if len(l1) != len(l2) {
		return false
	}
	if len(l1) == 0 {
		return true
	}

	for i, _ := range l1 {
		for l1[i] != nil {
			if l2[i] == nil {
				return false
			}
			if l1[i].Val != l2[i].Val {
				return false
			}
			l1[i] = l1[i].Next
			l2[i] = l2[i].Next
		}

		if l2[i] != nil {
			return false
		}
	}

	return true
}

func printList(l []*ListNode) (result [][]int) {
	for i, _ := range l {
		var r []int

		for l[i] != nil {
			r = append(r, l[i].Val)
			l[i] = l[i].Next
		}
		result = append(result, r)
	}

	return
}

func TestSplitListToParts(t *testing.T) {
	testCases := []struct {
		input  []int
		k      int
		output [][]int
	}{
		{[]int{1, 2, 3}, 5, [][]int{[]int{1}, []int{2}, []int{3}, []int{}, []int{}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7}, []int{8, 9, 10}}},
		{[]int{}, 3, [][]int{[]int{}, []int{}, []int{}}},
		{[]int{1, 2, 3}, 3, [][]int{[]int{1}, []int{2}, []int{3}}},
	}

	for _, tc := range testCases {
		head := buildList(tc.input)
		got := splitListToParts(head, tc.k)
		var want []*ListNode
		for _, a := range tc.output {
			h := buildList(a)
			want = append(want, h)
		}
		if !isSameList(got, want) {
			t.Errorf("wrong result, expected %v and got %v", printList(want), printList(got))
		}
	}
}
