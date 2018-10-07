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

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		input []int
		r     bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, null, 3, null, 3}, false},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, false},
	}

	for _, tc := range testCases {
		root := buildTree(tc.input)
		got1 := isSymmetric1(root)
		want := tc.r
		if got1 != want {
			t.Errorf("wrong result, expected %v and got %v", want, got1)
		}

		got2 := isSymmetric1(root)
		if got2 != want {
			t.Errorf("wrong result, expected %v and got %v", want, got2)
		}
	}
}
