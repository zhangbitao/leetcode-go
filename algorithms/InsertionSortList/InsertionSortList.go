// Source : https://leetcode-cn.com/problems/insertion-sort-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-10-04
//
// Sort a linked list using insertion sort.
//
// Algorithm of Insertion Sort:
//
// 1. Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
// 2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
// 3. It repeats until no input elements remain.

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var sorted, curr *ListNode
	var pp **ListNode
	for head != nil {
		curr = head
		head = head.Next

		pp = &sorted
		for *pp != nil && curr.Val > (*pp).Val {
			pp = &(*pp).Next
		}
		curr.Next = *pp
		*pp = curr
	}

	return sorted
}
