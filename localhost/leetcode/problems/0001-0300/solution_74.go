package solutions

/*
class Solution(object):

	def searchMatrix(self, matrix, target):
	    """
	    :type matrix: List[List[int]]
	    :type target: int
	    :rtype: bool
	    """
	    r = len(matrix)
	    c = len(matrix[0])

	    top, bot = 0, r - 1

	    while top <= bot:
	        row = (top+bot)//2

	        if target > matrix[row][-1]:
	            top = row + 1
	        elif target < matrix[row][0]:
	            bot = row - 1
	        else:
	            break

	    if not (top <= bot):
	        return False

	    low, high = 0, c - 1

	    row = (top+bot)//2

	    while low <= high:
	        mid = (low+high)//2
	        if target == matrix[row][mid]:
	            return True

	        if target > matrix[row][mid]:
	            low = mid + 1
	        else:
	            high = mid - 1

	    return False
*/
func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])

	top := 0
	bottom := r - 1

	for top <= bottom {
		row := (top + bottom) / 2

		if target > matrix[row][len(matrix[row])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			break
		}
	}

	if !(top <= bottom) {
		return false
	}

	row := (top + bottom) / 2

	low := 0
	high := c - 1

	for low <= high {
		mid := (low + high) / 2

		if matrix[row][mid] == target {
			return true
		}

		if target > matrix[row][mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
