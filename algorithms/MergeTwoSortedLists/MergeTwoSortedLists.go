// Source : https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-23
//
// Merge two sorted linked lists and return it as a new list. The new list should
// be made by splicing together the nodes of the first two lists.
//
// Example:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	fakeHead := &ListNode{
		Val:  0,
		Next: l1,
	}

	prev := fakeHead
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			prev = p1
			p1 = p1.Next
		} else {
			prev.Next = p2
			p2 = p2.Next
			prev = prev.Next
			prev.Next = p1
		}
	}

	if p2 != nil {
		prev.Next = p2
	}
	return fakeHead.Next
}
