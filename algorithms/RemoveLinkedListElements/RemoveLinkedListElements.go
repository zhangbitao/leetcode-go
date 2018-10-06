// Source : https://leetcode-cn.com/problems/remove-linked-list-elements/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-01
//
// Remove all elements from a linked list of integers that have value val.
//
// Example:
//
// Input:  1->2->6->3->4->5->6, val = 6
// Output: 1->2->3->4->5

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	fake := &ListNode{
		Next: head,
	}

	p := fake
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return fake.Next
}
