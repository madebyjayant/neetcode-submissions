func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix)-1, len(matrix[0])-1

	l,r := 0, cols

	for l<=rows && r>=0 {
		if matrix[l][r]<target {
			l++
			continue
		}
		if matrix[l][r]>target {
			r--
			continue
		}
		return true
	}
	return false
}
