// Source : https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-23
//
// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example:
//
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
//
// Output: 1->1->2->3->4->4->5->6

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var l0 *ListNode
	for i, l := range lists {
		if i == 0 {
			l0 = l
			continue
		}

		l0 = mergeTwoLists(l0, l)
	}
	return l0
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
