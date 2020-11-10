package example

/**
 * 数字范围按位与
 * 既然是范围，那除了右边的公共前缀，左边肯定是有1，0，在按位与的过程中会消失。最终保留的就是公共前缀
 * 因此，我们可以将问题重新表述为：给定两个整数，要求我们找到她们二进制字符串的公共前缀。
 *
 * 算法：
 * 如上述所说，算法由两个步骤组成：
 *
 * 我们通过右移，将两个数字压缩为它们的公共前缀。在迭代过程中，我们计算执行的右移操作数。
 * 将得到的公共前缀左移相同的操作数得到结果。
 *
 * 复杂度分析
 *
 * 时间复杂度：\mathcal{O}(1)O(1)。虽然算法中有一个循环，但是迭代次数是由整数的位数限定的，所以迭代次数是固定的。因此，算法的时间复杂度是常数级别的。
 * 空间复杂度：\mathcal{O}(1)O(1)，不管输入是什么，我们的内存消耗是常数的。
 **/
func rangeBitwiseAnd(m int, n int) int {
	var shift int = 0
	//find the common 1-bits
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}

/**
 * Brian Kernighan 算法
 * 说到位移，还有一个相关的算法叫做 Brian Kernighan 算法，它用于清除二进制串中最右边的 1。
 *
 * Brian Kernighan 算法的秘诀总结如下：
 * 当我们在 number 和 number-1 之间进行位运算时，原始 number 中最右边的 1 将变为 0。
 *
 * 基于上述技巧，我们可以应用它来计算两个位字符串的公共前缀。
 * 其思想是，对于给定的范围 [m，n]（即 m<n），我们可以对数字 n 迭代的应用技巧，清除最右边的 1，得到的数字我们表示为 n′。
 * 直到它小于或等于 m，最后，我们在 n′和 m 之间进行操作以获得最终结果。
 *
 * 通过应用 Brian Kernighan 算法，我们基本上清除了公共前缀右侧的位
 *
 */
func rangeBitwiseAnd_BrianKernighan(m int, n int) int {
	for m < n {
		//# turn off rightmost 1-bit
		n = n & (n - 1)
	}
	return m & n
}
