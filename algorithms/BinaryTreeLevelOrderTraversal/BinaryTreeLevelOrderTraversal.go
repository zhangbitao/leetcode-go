// Source : https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-08
//
// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
//
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	vv := [][]int{}

	for len(queue) > 0 {
		q := []*TreeNode{}
		v := []int{}
		for i := range queue {
			v = append(v, queue[i].Val)
			if queue[i].Left != nil {
				q = append(q, queue[i].Left)
			}
			if queue[i].Right != nil {
				q = append(q, queue[i].Right)
			}
		}
		vv = append(vv, v)
		queue = q
	}

	return vv
}

func levelOrder2(root *TreeNode) [][]int {
	vv := [][]int{}

	dfs(root, 0, &vv)
	return vv
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}
