// Source : https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-07
//
// Given a binary tree, return the inorder traversal of its nodes' values.
//
// Example:
//
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
//
// Output: [1,3,2]
// Follow up: Recursive solution is trivial, could you do it iteratively?

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var v []int

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			if len(stack) > 0 {
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				v = append(v, root.Val)
				root = root.Right
			}
		}
	}

	return v
}
