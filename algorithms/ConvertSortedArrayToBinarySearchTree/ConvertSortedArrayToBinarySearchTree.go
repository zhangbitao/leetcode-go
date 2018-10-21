// Source : https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-21
//
// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
// Example:
//
// Given the sorted array: [-10,-3,0,5,9],
//
// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//       0
//      / \
//    -3   9
//    /   /
//  -10  5

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	if mid == len(nums)-1 {
		node.Right = nil
	} else {
		node.Right = sortedArrayToBST(nums[mid+1:])
	}
	return node
}
