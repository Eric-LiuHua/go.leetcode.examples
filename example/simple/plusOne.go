package simple

//加一
func plusOne(digits []int) []int {
	var l int = len(digits)
	var tmp int = 1
	for i := l - 1; i >= 0; i-- {
		var cur = digits[i] + tmp
		if cur >= 10 {
			digits[i] = cur % 10
			tmp = cur / 10
		} else {
			digits[i] = cur
			tmp = 0
			break
		}
	}
	if tmp > 0 {
		var res []int = make([]int, 0)
		res = append(res, tmp)
		res = append(res, digits...)
		return res
	} else {
		return digits
	}
}
