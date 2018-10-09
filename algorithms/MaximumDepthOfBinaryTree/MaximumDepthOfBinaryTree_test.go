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

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		input []int
		depth int
	}{
		{[]int{3, 9, 20, null, null, 15, 7}, 3},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, null, 2}, 2},
		{[]int{1, null, 2, null, 3}, 3},
	}

	for _, tc := range testCases {
		root := buildTree(tc.input)
		got := maxDepth(root)
		want := tc.depth
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
