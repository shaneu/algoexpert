package sortedsquaredarray

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 5, 6, 8, 9},
			want:  []int{1, 4, 9, 25, 36, 64, 81},
		},
		{
			input: []int{-7, -5, -4, 3, 6, 8, 9},
			want:  []int{9, 16, 25, 36, 49, 64, 81},
		},
		{
			input: []int{-4, -2, 1, 3},
			want:  []int{1, 4, 9, 16},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := SortedSquaredArray(tt.input)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("SortedSquaredArray(%v), want %v; (-want, +got)\n%s", tt.input, tt.want, diff)
			}
		})
	}
}
