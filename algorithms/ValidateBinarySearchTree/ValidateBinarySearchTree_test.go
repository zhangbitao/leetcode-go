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

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		input  []int
		result bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{5, 1, 4, null, null, 3, 6}, false},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, null, 2}, true},
		{[]int{1, 2}, false},
	}

	for _, tc := range testCases {
		root := buildTree(tc.input)
		got1 := isValidBST1(root)
		if got1 != tc.result {
			t.Errorf("wrong result, expected %v and got %v", tc.result, got1)
		}

		got2 := isValidBST2(root)
		if got2 != tc.result {
			t.Errorf("wrong result, expected %v and got %v", tc.result, got2)
		}
	}
}
