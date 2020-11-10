package example

//二进制找唯一数，异或的特性
func OneTimesNumber(nums []int) int {
	var res int = 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
