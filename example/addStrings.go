package example

import "strconv"

/**
* 415. 字符串相加
* 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
* 注意：
* num1 和num2 的长度都小于 5100.
* num1 和num2 都只包含数字 0-9.
* num1 和num2 都不包含任何前导零。
* 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
**/
func addStrings(num1 string, num2 string) string {
	tmp := 0
	res := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || tmp > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmpvalue := tmp + x + y
		res = strconv.Itoa(tmpvalue%10) + res
		tmp = tmpvalue / 10
	}

	return res
}
