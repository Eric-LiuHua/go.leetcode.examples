package example

/**
* 771. 宝石与石头
**/
func numJewelsInStones(J string, S string) int {
	var res int = 0
	htable := make(map[rune]bool)
	for _, v := range J {
		htable[v] = true
	}

	for _, v := range S {
		if _, ok := htable[v]; ok {
			res++
		}
	}
	return res
}
