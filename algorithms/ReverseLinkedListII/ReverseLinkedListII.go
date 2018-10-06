// Source : https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-02
//
// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}

	fake := &ListNode{
		Next: head,
	}

	tail := fake
	for i := 0; i < m-1; i++ {
		if tail != nil {
			tail = tail.Next
		}
	}

	if tail == nil || tail.Next == nil {
		return head
	}

	p := tail.Next
	for i := m; i < n; i++ {
		if p.Next == nil {
			break
		}
		n := p.Next
		p.Next = n.Next
		n.Next = tail.Next
		tail.Next = n
	}

	return fake.Next
}
