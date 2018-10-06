// Source : https://leetcode-cn.com/problems/linked-list-components/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-10-06
//

package algorithms

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	if head == nil || len(G) == 0 {
		return 0
	}
	m := make(map[int]bool, 10000)
	for _, v := range G {
		m[v] = true
	}

	cnt := 0
	len := 0
	for p := head; p != nil; p = p.Next {
		if m[p.Val] == true {
			len++
		} else {
			if len > 0 {
				cnt++
				len = 0
			}
		}
	}
	if len > 0 {
		cnt++
	}

	return cnt
}
