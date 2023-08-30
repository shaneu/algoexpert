package twonumbersum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var cmpOpts = cmp.Options{
	cmpopts.SortSlices(func(a, b int) bool { return a < b }),
}

func TestTwoNumberSum(t *testing.T) {
	tests := []struct {
		array     []int
		targetSum int
		want      []int
	}{
		{
			array:     []int{3, 5, -4, 8, 11, 1, -1, 6},
			targetSum: 10,
			want:      []int{-1, 11},
		},
		{
			array:     []int{1},
			targetSum: 99,
			want:      []int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := TwoNumberSum(tt.array, tt.targetSum)

			if diff := cmp.Diff(tt.want, got, cmpOpts...); diff != "" {
				t.Errorf("TwoNumberSum(%v, %d), want %v; (-want, +got)\n%s", tt.array, tt.targetSum, tt.want, diff)
			}
		})
	}
}
