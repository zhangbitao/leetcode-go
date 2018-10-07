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

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		input1 []int
		input2 []int
		r      bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2}, []int{1, null, 2}, false},
		{[]int{1, 2, 1}, []int{1, 1, 2}, false},
		{[]int{}, []int{}, true},
		{[]int{1}, []int{}, false},
		{[]int{}, []int{1}, false},
	}

	for _, tc := range testCases {
		p := buildTree(tc.input1)
		q := buildTree(tc.input2)
		got := isSameTree(p, q)
		want := tc.r
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
