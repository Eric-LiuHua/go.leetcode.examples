package example

/**
 * 461. 汉明距离
 * 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
 * 要求两个数字对应二进制位不同的位置的数目，异或运算:相同为1，不同为0；
 * 那么汉明距离即为：将两个数字异或得到的数，计算其二进制表示的1的个数。
 *
 */
func hammingDistance(x int, y int) int {
	var res int = 0
	var n = x ^ y
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return res
}
