package nonconstructiblechange

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	change := 0

	for _, coin := range coins {
		if coin > change+1 {
			break
		}
		change += coin
	}

	return change + 1
}
