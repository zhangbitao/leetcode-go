// Source : https://leetcode-cn.com/problems/add-two-numbers/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-09-24
//
// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order and each of their nodes
// contain a single digit. Add the two numbers and return it as a linked
// list.
//
// You may assume the two numbers do not contain any leading zero, except
// the number 0 itself.
//
// Example:
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode
    r := &result
    remainder := 0
    for l1 != nil || l2 != nil {   
        var v1, v2 int
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        n := &ListNode{
            Val: (remainder + v1 + v2) % 10,
        }
        *r = n
        r = &(n.Next)
        
		remainder = (remainder + v1 + v2) / 10

        if remainder > 0 {
            n := &ListNode{
                Val: remainder,
            }
            *r = n
        }
    }
    return result
}