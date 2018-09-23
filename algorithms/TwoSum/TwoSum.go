// Source : https://leetcode-cn.com/problems/two-sum/description/
// Author : zhangbitao <xszhangbitao@gmail.com>
// Date   : 2018-09-24
//
// Given an array of integers, return indices of the two numbers such that
// they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you
// may not use the same element twice.
//
// Example:
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

package algorithms

func TwoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)

	for i, num := range nums {
		if v, ok := m[num]; ok {
			result = append(result, v)
			result = append(result, i)
			break
		} else {
			m[target-num] = i
		}
	}

	return result
}
