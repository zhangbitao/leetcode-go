package algorithms

import (
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	testCases := []struct {
		n      int
		output int
	}{
		{3, 5},
		{0, 0},
		{1, 1},
		{2, 2},
	}

	for _, tc := range testCases {
		got := numTrees(tc.n)
		want := tc.output
		if got != want {
			t.Errorf("wrong result, expected %v and got %v", want, got)
		}
	}
}
