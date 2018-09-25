// Source : https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-25
//
// Given a sorted linked list, delete all duplicates such that each element appear only once.
//
// Example 1:
//
// Input: 1->1->2
// Output: 1->2
// Example 2:
//
// Input: 1->1->2->3->3
// Output: 1->2->3

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return head
}
