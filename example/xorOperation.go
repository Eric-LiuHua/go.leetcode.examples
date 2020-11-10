package example

/**
* 1486. 数组异或操作
* 给你两个整数，n 和 start 。
* 数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。
* 请返回 nums 中所有元素按位异或（XOR）后得到的结果。
*
* 少一次循环，内存使用更少 ，1.9m
*
 */
func xorOperation(n int, start int) int {
	var res int = start
	for i := 1; i < n; i++ {
		res ^= start + i*2
	}
	return res
}

//多一次循环，话费2m
func xorOperation1(n int, start int) int {
	var res int = 0
	for i := 0; i < n; i++ {
		res ^= start + i*2
	}
	return res
}
