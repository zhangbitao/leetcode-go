// Source : https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-10-02
//
// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
//
// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
// Example:
//
// Given the sorted linked list: [-10,-3,0,5,9],
//
// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//       0
//      / \
//    -3   9
//    /   /
//  -10  5

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	len := 0
	for p := head; p != nil; p = p.Next {
		len++
	}

	return sortedListToBSTHelper(head, 0, len-1)
}

func sortedListToBSTHelper(head *ListNode, low, high int) *TreeNode {
	if low > high || head == nil {
		return nil
	}

	mid := low + (high-low)/2
	left := sortedListToBSTHelper(head, low, mid-1)

	p := head
	for i := 0; i < mid; i++ {
		p = p.Next
	}

	node := &TreeNode{}
	node.Val = p.Val
	node.Left = left

	right := sortedListToBSTHelper(head, mid+1, high)
	node.Right = right
	return node
}
