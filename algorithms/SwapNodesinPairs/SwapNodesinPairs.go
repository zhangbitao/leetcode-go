// Source : https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-09-24
//
// Given a linked list, swap every two adjacent nodes and return its head.
//
// Example:
//
// Given 1->2->3->4, you should return the list as 2->1->4->3.
//
// Note:
//
// Your algorithm should use only constant extra space.
//You may not modify the values in the list's nodes, only nodes itself may be changed.

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var s *ListNode

	for p := head; p != nil && p.Next != nil; p = p.Next {
		n := p.Next
		p.Next = n.Next
		n.Next = p

		if s != nil {
			s.Next = n
		}
		s = p

		if p == head {
			head = n
		}
	}
	return head
}
