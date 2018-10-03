package algorithms

import (
	"container/list"
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

func bstToArray(root *TreeNode) (r []int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)
	for q.Len() != 0 {
		e := q.Front()
		node := e.Value.(*TreeNode)

		r = append(r, node.Val)
		if node.Left != nil {
			q.PushBack(node.Left)
		}
		if node.Right != nil {
			q.PushBack(node.Right)
		}
		q.Remove(e)
	}

	return r
}

func isSameArray(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i, _ := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func TestSortedListToBST(t *testing.T) {
	testCases := []struct {
		input1 []int
		output []int
	}{
		{[]int{-10, -3, 0, 5, 9}, []int{0, -10, 5, -3, 9}},
	}

	for _, tc := range testCases {
		head := buildList(tc.input1)
		g := sortedListToBST(head)
		got := bstToArray(g)
		want := tc.output
		if !isSameArray(got, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
