// Source : https://leetcode-cn.com/problems/validate-binary-search-tree/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-20
//
// Given a binary tree, determine if it is a valid binary search tree (BST).
//
// Assume a BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// Example 1:
//
// Input:
//     2
//    / \
//   1   3
// Output: true
//
// Example 2:
//
//     5
//    / \
//   1   4
//      / \
//     3   6
// Output: false
// Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
//              is 5 but its right child's value is 4.

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	stack := []*TreeNode{}
	node := root
	v := []int{}
	for len(stack) != 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(v) != 0 && node.Val <= v[len(v)-1] {
				return false
			}
			v = append(v, node.Val)
			node = node.Right
		}
	}
	return true
}

// depth-first search
func isValidBST2(root *TreeNode) bool {
	v := []int{}
	dfs(root, &v)
	for i := 0; i < len(v)-1; i++ {
		if v[i] >= v[i+1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, v *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, v)
	*v = append(*v, root.Val)
	dfs(root.Right, v)
}
