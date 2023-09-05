package validatesubsequence

import (
	"fmt"
	"testing"
)

func TestIsValidSubsequence(t *testing.T) {
	tests := []struct {
		array       []int
		subsequence []int
		want        bool
	}{
		{
			array:       []int{5, 1, 22, 25, 6, -1, 8, 10},
			subsequence: []int{1, 6, -1, 10},
			want:        true,
		},
		{
			array:       []int{5, 1, 22, 25, 6, -1, 8, 10},
			subsequence: []int{1, 6, -1, 11},
			want:        false,
		},
		{
			array:       []int{5},
			subsequence: []int{1},
			want:        false,
		},
		{
			array:       []int{5, 1, 22, 25, 6, -1, 8, 10},
			subsequence: []int{8},
			want:        true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {

			got := IsValidSubsequence(tt.array, tt.subsequence)
			if got != tt.want {
				t.Errorf("IsValidSubsequence(%v, %v): got %v, want %v", tt.array, tt.subsequence, got, tt.want)
			}
		})
	}
}
