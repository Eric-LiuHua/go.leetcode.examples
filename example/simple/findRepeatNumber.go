package simple

//剑指 Offer 03. 数组中重复的数字,限制：2 <= n <= 100000
func findRepeatNumber(nums []int) int {
	var res int = -1
	var iset map[int]bool = make(map[int]bool)
	for _, num := range nums {
		_, ok := iset[num]
		if ok {
			res = num
			break
		} else {
			iset[num] = true
		}
	}
	return res
}
