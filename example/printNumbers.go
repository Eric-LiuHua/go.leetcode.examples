package example

/**
 * 剑指 Offer 17. 打印从1到最大的n位数
 * 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
 * 核心找到n位数的最大值。
 * 会有溢出的风险
 *
 */
func printNumbers(n int) []int {
	var res []int = make([]int, 0)
	var max int = 0
	for i := n; i > 0; i-- {
		max = max*10 + 9
	}
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}
