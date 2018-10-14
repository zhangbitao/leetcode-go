package algorithms

import (
	"math"
	"testing"
)

const (
	null = math.MinInt32
)

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

func TestGenerateTrees(t *testing.T) {
	testCases := []struct {
		n      int
		output [][]int
	}{
		{
			3,
			[][]int{
				[]int{1, null, 2, null, 3},
				[]int{1, null, 3, 2},
				[]int{2, 1, 3},
				[]int{3, 1, null, null, 2},
				[]int{3, 2, null, 1},
			},
		},
		{
			0,
			[][]int{},
		},
		{
			1,
			[][]int{
				[]int{1},
			},
		},
		{
			2,
			[][]int{
				[]int{1, null, 2},
				[]int{2, 1},
			},
		},
	}

	for _, tc := range testCases {
		trees := generateTrees(tc.n)

		var got [][]int
		for _, root := range trees {
			got = append(got, bstToArray(root))
		}
		want := tc.output
		if !isSameArray(got, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
