package transposematrix

func TransposeMatrix(matrix [][]int) [][]int {
	var out [][]int

	for columnIdx := range matrix[0] {
		var newRow []int

		for rowIdx := range matrix {
			newRow = append(newRow, matrix[rowIdx][columnIdx])
		}

		out = append(out, newRow)
	}

	return out
}
