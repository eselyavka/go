package solutions

func maxEqualRowsAfterFlips(matrix [][]int) int {
	ans := 0
	rowsHash := make(map[string]int)

	for _, row := range matrix {
		itemsSet := make(map[int]struct{})
		for _, col := range row {
			itemsSet[col] = struct{}{}
		}
		if len(itemsSet) == 1 {
			ans++
		}
		rowsHash[sliceToString(row)]++
	}

	for _, row := range matrix {
		xor := sliceToString(flipArrayXOR(row))
		ans = MaxInts([]int{ans, rowsHash[sliceToString(row)] + rowsHash[xor]})
	}

	return ans
}
