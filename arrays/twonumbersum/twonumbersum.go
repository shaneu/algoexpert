package twonumbersum

func TwoNumberSum(array []int, target int) []int {
	need := make(map[int]struct{})

	for _, x := range array {
		want := target - x

		_, ok := need[want]
		if ok {
			return []int{x, want}
		}

		need[x] = struct{}{}
	}

	return make([]int, 0)
}
