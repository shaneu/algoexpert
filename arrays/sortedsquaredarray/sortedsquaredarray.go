package sortedsquaredarray

import "math"

func SortedSquaredArray(array []int) []int {
	out := make([]int, len(array))

	start := 0
	end := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		if math.Abs(float64(array[start])) > math.Abs(float64(array[end])) {
			out[i] = array[start] * array[start]
			start++
		} else {
			out[i] = array[end] * array[end]
			end--
		}
	}

	return out
}
