// Source : https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-23
//
// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.
//
// Note:
//
// Given n will always be valid.

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return nil
	}
	fakeHead := ListNode{
		Val:  0,
		Next: head,
	}
	head = &fakeHead
	p1 := head
	p2 := head
	for i := 0; i < n; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p1.Next = p1.Next.Next

	return head.Next
}
