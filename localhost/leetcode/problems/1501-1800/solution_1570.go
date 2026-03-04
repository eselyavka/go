package solutions

func ConstructorSparseVector(nums []int) SparseVector {
	vector := SparseVector{}
	vector.storage = make(map[int]int)

	for idx, val := range nums {
		if val != 0 {
			vector.storage[idx] = val
		}

	}
	return vector
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	res := 0
	for idx, val1 := range this.storage {
		if val2, found := vec.storage[idx]; found {
			res += val1 * val2
		}
	}
	return res
}
