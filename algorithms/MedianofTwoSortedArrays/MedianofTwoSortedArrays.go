// Source : https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
// Author : zhangbitao <zhangbitao01@gmail.com>
// Date   : 2018-09-24
//
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Find the median of the two sorted arrays. The overall run time complexity
// should be O(log (m+n)).
//
// Example 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// The median is 2.0
// Example 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// The median is (2 + 3)/2 = 2.5

package algorithms

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	m := n1 + n2
	if m%2 == 1 {
		return (float64)(findMedianSortedArraysHelper(nums1, n1, nums2, n2, m/2+1))
	} else {
		return (float64)(findMedianSortedArraysHelper(nums1, n1, nums2, n2, m/2)+
			findMedianSortedArraysHelper(nums1, n1, nums2, n2, m/2+1)) / 2
	}
}

func findMedianSortedArraysHelper(nums1 []int, n1 int, nums2 []int, n2 int, m int) int {
	if n1 > n2 {
		return findMedianSortedArraysHelper(nums2, n2, nums1, n1, m)
	}

	if n1 == 0 {
		return nums2[m-1]
	}
	if m == 1 {
		return min(nums1[0], nums2[0])
	}

	index1 := min(m/2, n1)
	index2 := m - index1
	if nums1[index1-1] < nums2[index2-1] {
		return findMedianSortedArraysHelper(nums1[index1:], n1-index1, nums2, n2, m-index1)
	} else if nums1[index1-1] > nums2[index2-1] {
		return findMedianSortedArraysHelper(nums1, n1, nums2[index2:], n2-index2, m-index2)
	} else {
		return nums1[index1-1]
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
