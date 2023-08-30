package twonumbersum

func TwoNumberSum(array []int, target int) []int {
	nums := make(map[int]struct{})

	for _, x := range array {
		want := target - x

		_, ok := nums[want]
		if ok {
			return []int{x, want}
		}

		nums[x] = struct{}{}
	}

	return make([]int, 0)
}
