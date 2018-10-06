// Source : https://leetcode-cn.com/problems/partition-list/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-01
//
// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
// You should preserve the original relative order of the nodes in each of the two partitions.
//
// Example:
//
// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	fake := &ListNode{
		Next: head,
	}

	pos := fake
	p := fake
	for p.Next != nil {
		if p.Next.Val < x {
			if p == pos {
				pos = pos.Next
				p = p.Next
			} else {
				t := p.Next
				p.Next = t.Next
				t.Next = pos.Next
				pos.Next = t
				pos = t
			}
		} else {
			p = p.Next
		}
	}
	return fake.Next
}
