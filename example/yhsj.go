package example

//打印杨辉三角
func YhsjA(rowIndex int) [][]int {
	var res [][]int = make([][]int, rowIndex)
	res[0] = make([]int, 1)
	res[0][0] = 1

	for i := 1; i < rowIndex; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i][i] = 1
	}
	return res
}

//打印杨辉三角某一行,这里指的是下标行。所以需要加一行。
func YhsjB(rowIndex int) []int {
	rowIndex++
	return YhsjA(rowIndex)[rowIndex-1]
}
