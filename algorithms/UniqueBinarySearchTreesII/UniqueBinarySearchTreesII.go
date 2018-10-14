// Source : https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-14
//
// Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
//
// Example:
//
// Input: 3
// Output:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// Explanation:
// The above output corresponds to the 5 unique BST's shown below:

//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package algorithms

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateTreesHelper(1, n)
}

func generateTreesHelper(low int, high int) (v []*TreeNode) {
	if low > high || low <= 0 || high <= 0 {
		return []*TreeNode{nil}
	}

	if low == high {
		return []*TreeNode{&TreeNode{Val: low}}
	}

	for i := low; i <= high; i++ {
		vleft := generateTreesHelper(low, i-1)
		vright := generateTreesHelper(i+1, high)
		for l := 0; l < len(vleft); l++ {
			for r := 0; r < len(vright); r++ {
				root := &TreeNode{Val: i}
				root.Left = vleft[l]
				root.Right = vright[r]
				v = append(v, root)
			}
		}
	}
	return
}
