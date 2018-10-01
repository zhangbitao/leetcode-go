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

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input []int
		ok    bool
	}{
		{[]int{1, 2}, false},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 2, 2}, false},
		{[]int{1, 2, 3, 2, 1}, true},

		{[]int{1, 2, 3, 4, 2, 1}, false},
		{[]int{1}, true},
		{[]int{}, true},
	}

	for _, tc := range testCases {
		head := buildList(tc.input)
		got := isPalindrome(head)
		want := tc.ok
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
