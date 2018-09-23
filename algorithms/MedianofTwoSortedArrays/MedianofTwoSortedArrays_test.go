package algorithms

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		nums1  []int
		nums2  []int
		result float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1}, []int{2, 3, 4}, 2.5},
		{[]int{}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, []int{}, 2.5},
		{[]int{100001}, []int{100000}, 100000.5},
		{[]int{1, 11, 22, 33, 44, 55}, []int{2, 12, 23, 34, 45, 56, 60}, 33},
		{[]int{1, 11, 22, 33, 44, 55}, []int{2, 12, 23, 34, 45, 56}, 28},
		{[]int{1, 2, 5, 6, 8}, []int{11, 13, 17, 30, 45, 50}, 11},
	}

	for _, tc := range testCases {
		got := findMedianSortedArrays(tc.nums1, tc.nums2)
		want := tc.result
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
