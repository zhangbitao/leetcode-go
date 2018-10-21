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

func bstToArray(root *TreeNode) (r []int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	r = append(r, root.Val)

	for len(queue) > 0 {
		q := []*TreeNode{}
		for i := range queue {
			if queue[i].Left == nil && queue[i].Right == nil {
				continue
			}
			if queue[i].Left != nil {
				q = append(q, queue[i].Left)
				r = append(r, queue[i].Left.Val)
			} else {
				r = append(r, null)
			}

			if queue[i].Right != nil {
				q = append(q, queue[i].Right)
				r = append(r, queue[i].Right.Val)
			} else {
				r = append(r, null)
			}
		}
		queue = q
	}

	index := len(r)
	for index >= 0 {
		if r[index-1] != null {
			break
		}
		index--
	}
	r = r[:index]

	return
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

func TestSortedArrayToBST(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{-10, -3, 0, 5, 9}, []int{0, -3, 9, -10, null, 5}},
		{[]int{}, []int{}},
		{[]int{-10}, []int{-10}},
		{[]int{-10, -3, 0}, []int{-3, -10, 0}},
		{[]int{-10, -3, 0, 5}, []int{0, -3, 5, -10}},
	}

	for _, tc := range testCases {
		root := sortedArrayToBST(tc.input)
		got := bstToArray(root)
		want := tc.output
		if !isSameArray(got, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
