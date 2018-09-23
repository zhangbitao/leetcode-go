package algorithms

import (
	"testing"
)

func isSame(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{15, 11, 7, 2}, 17, []int{0, 3}},
		{[]int{2, 11, 7, 2}, 4, []int{0, 3}},
	}

	for _, tc := range testCases {
		got := TwoSum(tc.nums, tc.target)
		want := tc.result
		if !isSame(got, want) {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
