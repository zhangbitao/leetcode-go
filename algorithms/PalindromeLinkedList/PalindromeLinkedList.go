// Source : https://leetcode-cn.com/problems/palindrome-linked-list/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-10-01
//
// Given a singly linked list, determine if it is a palindrome.
//
// Example 1:
//
// Input: 1->2
// Output: false
// Example 2:
//
// Input: 1->2->2->1
// Output: true
//
// Follow up:
// Could you do it in O(n) time and O(1) space?

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head.Next
	p1 := slow
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = slow.Next.Next
		tmp.Next = p1
		p1 = tmp
	}

	p2 := slow.Next
	if fast.Next != nil {
		p2 = p2.Next
	}
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
