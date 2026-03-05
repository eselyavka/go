package p1072

import "localhost/leetcode/util"

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
		rowsHash[util.SliceToString(row)]++
	}

	for _, row := range matrix {
		xor := util.SliceToString(util.FlipArrayXOR(row))
		ans = util.MaxInts([]int{ans, rowsHash[util.SliceToString(row)] + rowsHash[xor]})
	}

	return ans
}
