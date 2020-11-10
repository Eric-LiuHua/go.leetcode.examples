package example

import (
	"fmt"
	"strings"
)

/**
 * 剑指 Offer 15. 二进制中1的个数
 *
 *
 *
 **/
func hammingWeight(num uint32) int {
	var res int = 0
	tmp := num
	for tmp != 0 {
		if (tmp & 1) == 1 {
			res++
		}
		tmp >>= 1
	}

	return res
}

/**
 * 剑指 Offer 15. 二进制中1的个数
 * %b 二进制
 *
 *
 **/
func hammingWeightA(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}
