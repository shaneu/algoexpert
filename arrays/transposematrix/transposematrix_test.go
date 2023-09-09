package transposematrix

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTransposeMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{
				{1, 2},
			},
			want: [][]int{
				{1},
				{2},
			},
		},
		{
			matrix: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tests:%d", i), func(t *testing.T) {
			got := TransposeMatrix(tt.matrix)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("TransposeMatrix(%v), want %v; (-want, +got)\n%s", tt.matrix, tt.want, diff)
			}
		})
	}
}
