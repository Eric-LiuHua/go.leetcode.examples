package simple

//有效的数独
func isValidSudoku(board [][]byte) bool {
	//init data
	var xc map[int]map[int]int = make(map[int]map[int]int)
	var yc map[int]map[int]int = make(map[int]map[int]int)
	var basec map[int]map[int]int = make(map[int]map[int]int)
	var x int = len(board)
	var y int = len(board[0])
	for i := 0; i < x; i++ {
		xc[i] = make(map[int]int)
		yc[i] = make(map[int]int)
		basec[i] = make(map[int]int)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			tmp := board[i][j]
			if tmp != '.' {
				tmpv := int(tmp)
				bpoint := (i/3)*(j/3) + (i / 3)
				if _, ok := xc[i][tmpv]; ok {
					return false
				} else {
					xc[i][tmpv] = 1
				}
				if _, ok := yc[j][tmpv]; ok {
					return false
				} else {
					yc[j][tmpv] = 1
				}
				if _, ok := basec[bpoint][tmpv]; ok {
					return false
				} else {
					basec[bpoint][tmpv] = 1
				}
			}
		}

	}
	return true
}
