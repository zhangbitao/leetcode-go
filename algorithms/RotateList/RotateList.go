// Source : https://leetcode-cn.com/problems/rotate-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-24
//
// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Example 1:
//
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL
//
// Example 2:
//
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}
	len := 1
	p := head
	for p.Next != nil {
		p = p.Next
		len++
	}

	p.Next = head
	k = len - k%len
	for i := 0; i < k; i++ {
		p = p.Next
	}

	head = p.Next
	p.Next = nil

	return head
}
