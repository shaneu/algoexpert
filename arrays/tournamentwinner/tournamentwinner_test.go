package tournamentwinner

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTournamentWinner(t *testing.T) {
	tests := []struct {
		competitions [][]string
		results      []int
		want         string
	}{
		{
			competitions: [][]string{
				{"HTML", "C#"},
				{"C#", "Python"},
				{"Python", "HTML"},
			},
			results: []int{0, 0, 1},
			want:    "Python",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := TournamentWinner(tt.competitions, tt.results)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("TournamentWinner(%v, %v), want %s; (-want, +got)\n%s", tt.competitions, tt.results, tt.want, diff)
			}
		})
	}
}
