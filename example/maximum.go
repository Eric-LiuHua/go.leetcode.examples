package example

//面试题 16.07. 最大数值
//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。
// a > b时，ret = a - 0
// a < b时，ret = a - (a - b)
//
// 通过位运算来决定 a 后面的值
// 表达式 res & (res >> 63)
// res >> 63，前面 已经说过，要么全为 1， 要么全为0
// 代入运算 res & (res >> 63),结果有两个
// 0 也就是 ret & (00000000 ... 00000000)
// ret 也就是 ret & (11111111 ... 11111111)
func maximum(a int, b int) int {
	//res 为ab的差额,由于边界值，需要换64位来存储
	res := int64(a - b)
	res = int64(a) - res&(res>>63)

	return int(res)
}
