package validatesubsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx := 0

	for _, val := range array {
		if seqIdx == len(sequence) {
			break
		}

		if val == sequence[seqIdx] {
			seqIdx++
		}
	}

	return seqIdx == len(sequence)
}
