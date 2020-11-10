package example

//杨辉三角
func PascalTriangle(level int) [][]int {
	var res [][]int = make([][]int, level)
	res[0] = make([]int, 1)
	res[0][0] = 1

	for r := 1; r <= level-1; r++ {
		res[r] = make([]int, r+1)
		res[r][0] = 1

		for c := 1; c < r; c++ {
			res[r][c] = res[r-1][c-1] + res[r-1][c]
		}

		res[r][r] = 1
	}

	return res
}
