// Source : https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-09-25
//
// Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
// Example 1:
//
// Input: 1->2->3->3->4->4->5
// Output: 1->2->5
// Example 2:
//
// Input: 1->1->1->2->3
// Output: 2->3

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

	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}

	dup := false
	prev := fakeHead
	for p := head; p != nil && p.Next != nil; p = p.Next {
		if dup == false && p.Val == p.Next.Val {
			dup = true
			continue
		}
		if dup == true && p.Val != p.Next.Val {
			dup = false
			prev.Next = p.Next
			continue
		}
		if dup == false {
			prev = p
		}
	}
	if dup == true {
		prev.Next = nil
	}
	return fakeHead.Next
}
