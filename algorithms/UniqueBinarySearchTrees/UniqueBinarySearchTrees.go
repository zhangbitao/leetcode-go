// Source : https://leetcode-cn.com/problems/unique-binary-search-trees/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-10-14
//
// Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
//
// Example:
//
// Input: 3
// Output: 5
// Explanation:
// Given n = 3, there are a total of 5 unique BST's:
//
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package algorithms

func numTrees(n int) int {
	if n <= 1 {
		return n
	}

	cnt := make([]int, n+1)
	cnt[0], cnt[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			cnt[i] += cnt[j] * cnt[i-j-1]
		}
	}
	return cnt[n]
}
