// Source : https://leetcode-cn.com/problems/reorder-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-10-03
//
// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example 1:
//
// Given 1->2->3->4, reorder it to 1->4->2->3.
//
// Example 2:
//
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var p *ListNode
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	if p != nil {
		p.Next = nil
	}

	l1 := head
	l2 := reverseList(p1)

	head = mergeList(l1, l2)
}

func reverseList(head *ListNode) *ListNode {
	var h, p *ListNode
	for head != nil {
		p = head
		head = head.Next
		p.Next = h
		h = p
	}
	return h
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	var n1, n2 *ListNode
	for p1 != nil && p2 != nil {
		n1 = p1.Next
		n2 = p2.Next

		p1.Next = p2
		p2.Next = n1
		if n1 == nil {
			p2.Next = n2
			break
		}

		p1 = n1
		p2 = n2
	}
	return p1
}
