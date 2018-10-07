package algorithms

import (
	"math"
	"testing"
)

const (
	null = math.MinInt32
)

func buildTree(value []int) *TreeNode {
	if len(value) <= 0 {
		return nil
	}

	tree := make([]*TreeNode, len(value))
	for i, v := range value {
		if v == null {
			tree[i] = nil
		} else {
			tree[i] = &TreeNode{
				Val: v,
			}
		}
	}

	pos := 1
	n := len(value)
	for i := 0; i < n && pos < n; i++ {
		if tree[i] != nil {
			tree[i].Left = tree[pos]
			pos++
			if pos < n {
				tree[i].Right = tree[pos]
				pos++
			}
		}
	}
	return tree[0]
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

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, null, 2, 3}, []int{1, 3, 2}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 2, 5, 1, 6, 3}},
	}

	for _, tc := range testCases {
		root := buildTree(tc.input)
		got := inorderTraversal(root)
		want := tc.output
		if !isSameArray(got, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
