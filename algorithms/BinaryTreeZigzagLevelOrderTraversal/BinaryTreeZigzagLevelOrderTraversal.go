// Source : https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-09
//
// Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
//
// return its zigzag level order traversal as:
//
// [
//   [3],
//   [20,9],
//   [15,7]
// ]

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	vv := [][]int{}
	zigzag := false

	for len(queue) > 0 {
		q := []*TreeNode{}
		v := []int{}
		for i := range queue {
			if zigzag {
				v = append(v, queue[len(queue)-1-i].Val)
			} else {
				v = append(v, queue[i].Val)
			}
			if queue[i].Left != nil {
				q = append(q, queue[i].Left)
			}
			if queue[i].Right != nil {
				q = append(q, queue[i].Right)
			}
		}
		zigzag = !zigzag
		vv = append(vv, v)
		queue = q
	}

	return vv
}
