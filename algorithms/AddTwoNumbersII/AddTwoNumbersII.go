// Source : https://leetcode-cn.com/problems/add-two-numbers-ii/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-04
//
// You are given two non-empty linked lists representing two non-negative integers. The most significant
// digit comes first and each of their nodes contain a single digit. Add the two numbers and return it
// as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Follow up:
// What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
//
// Example:
//
// Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 8 -> 0 -> 7

package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := getLength(l1)
	len2 := getLength(l2)

	fake := &ListNode{
		Val: 1,
	}
	if len1 > len2 {
		fake.Next = addTwoNumbersHelper(l1, l2, len1-len2)
	} else {
		fake.Next = addTwoNumbersHelper(l2, l1, len2-len1)
	}

	if fake.Next.Val > 9 {
		fake.Next.Val = fake.Next.Val % 10
		return fake
	}

	return fake.Next
}

func getLength(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, diff int) *ListNode {
	if l1 == nil {
		return nil
	}

	var res, pos *ListNode
	if diff == 0 {
		res = &ListNode{
			Val: l1.Val + l2.Val,
		}
		pos = addTwoNumbersHelper(l1.Next, l2.Next, 0)
	} else {
		res = &ListNode{
			Val: l1.Val,
		}
		pos = addTwoNumbersHelper(l1.Next, l2, diff-1)
	}

	if pos != nil && pos.Val > 9 {
		pos.Val = pos.Val % 10
		res.Val++
	}
	res.Next = pos

	return res
}
