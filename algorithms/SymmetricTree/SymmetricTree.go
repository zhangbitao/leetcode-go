// Source : https://leetcode-cn.com/problems/symmetric-tree/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-07
//
// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//
// But the following [1,2,2,null,3,null,3] is not:
//
//     1
//    / \
//   2   2
//    \   \
//    3    3
//
// Note:
// Bonus points if you could solve it both recursively and iteratively.

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursively
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree1(root.Left, root.Right)
}

func isSymmetricTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSymmetricTree1(p.Left, q.Right) && isSymmetricTree1(p.Right, q.Left)
}

// iteratively
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree2(root.Left, root.Right)
}

func isSymmetricTree2(p *TreeNode, q *TreeNode) bool {
	var queue1, queue2 []*TreeNode

	queue1 = append(queue1, p)
	queue2 = append(queue2, q)
	for len(queue1) > 0 && len(queue2) > 0 {
		q1 := queue1[0]
		queue1 = queue1[1:]
		q2 := queue2[0]
		queue2 = queue2[1:]

		if q1 == nil && q2 == nil {
			continue
		}
		if q1 == nil || q2 == nil {
			return false
		}
		if q1.Val != q2.Val {
			return false
		}

		queue1 = append(queue1, q1.Left)
		queue2 = append(queue2, q2.Right)

		queue1 = append(queue1, q1.Right)
		queue2 = append(queue2, q2.Left)
	}
	return true
}
