// Source : https://leetcode-cn.com/problems/sort-list/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-04
//
// Sort a linked list in O(n log n) time using constant space complexity.
//
// Example 1:
//
// Input: 4->2->1->3
// Output: 1->2->3->4
//
// Example 2:
//
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	p2 = p1.Next
	p1.Next = nil
	return mergeLists(sortList(head), sortList(p2))
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fake := &ListNode{}
	p := fake

	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return fake.Next
}
