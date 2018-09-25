// Source : https://leetcode-cn.com/problems/linked-list-cycle/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-25
//
// Given a linked list, determine if it has a cycle in it.
//
// Follow up:
// Can you solve it without using extra space?

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
	}
	return false
}
