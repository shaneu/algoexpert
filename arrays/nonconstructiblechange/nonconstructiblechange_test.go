package nonconstructiblechange

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNonConstructibleChange(t *testing.T) {
	tests := []struct {
		coins []int
		want  int
	}{
		{
			coins: []int{1, 2, 5},
			want:  4,
		},
		{
			coins: []int{5, 7, 1, 1, 2, 3, 22},
			want:  20,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := NonConstructibleChange(tt.coins)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NonConstructibleChange(%v), want %d; (-want, +got)\n%s", tt.coins, tt.want, diff)
			}
		})
	}
}
