// Source : https://leetcode-cn.com/problems/reverse-linked-list/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-01
//
// Reverse a singly linked list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
//
// Follow up:
//
// A linked list can be reversed either iteratively or recursively. Could you implement both?

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var h, p *ListNode
	for head != nil {
		p = head.Next
		head.Next = h
		h = head
		head = p
	}
	return h
}
