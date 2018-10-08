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

func isSameArray(a1 [][]int, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i, _ := range a1 {
		if len(a1[i]) != len(a2[i]) {
			return false
		}
		for j, _ := range a1[i] {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		input  []int
		output [][]int
	}{
		{
			[]int{3, 9, 20, null, null, 15, 7},
			[][]int{[]int{3}, []int{9, 20}, []int{15, 7}},
		},
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{1},
			[][]int{[]int{1}},
		},
		{
			[]int{1, 2},
			[][]int{[]int{1}, []int{2}},
		},
		{
			[]int{1, null, 2},
			[][]int{[]int{1}, []int{2}},
		},
		{
			[]int{1, null, 2, null, 3},
			[][]int{[]int{1}, []int{2}, []int{3}},
		},
	}

	for _, tc := range testCases {
		root := buildTree(tc.input)
		got1 := levelOrder1(root)
		want := tc.output
		if !isSameArray(got1, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got1)
		}

		got2 := levelOrder2(root)
		if !isSameArray(got2, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got2)
		}

	}
}
