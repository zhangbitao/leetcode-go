// Source : https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-24
//
// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
// k is a positive integer and is less than or equal to the length of the linked list. If the number
// of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
// Example:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be changed.

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := fakeHead

	for p != nil {
		p.Next = reverseList(p.Next, k)
		for i := 0; i < k && p != nil; i++ {
			p = p.Next
		}
	}

	return fakeHead.Next
}

func reverseList(head *ListNode, k int) *ListNode {
	end := head
	for end != nil && k > 0 {
		end = end.Next
		k--
	}
	if k > 0 {
		return head
	}

	start := end
	p := head
	for p != end {
		q := p.Next
		p.Next = start
		start = p
		p = q
	}
	return start
}
