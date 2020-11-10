package dynamic_programming

//474. 一和零
//动态规划
//这道题和经典的背包问题很类似，不同的是在背包问题中，
//我们只有一种容量，而在这道题中，我们有 0 和 1 两种容量。
//每个物品（字符串）需要分别占用 0 和 1 的若干容量，并且所有物品的价值均为 1。因此我们可以使用二维的动态规划。
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zero, one := count(s)
		for mi := m; mi >= zero; mi-- {
			for ni := n; ni >= one; ni-- {
				if dp[mi][ni] < (1 + dp[mi-zero][ni-one]) {
					dp[mi][ni] = (1 + dp[mi-zero][ni-one])
				}
			}
		}
	}
	return dp[m][n]
}

//统计0 1 的数量
func count(str string) (zero, one int) {
	zero, one = 0, 0
	for _, s := range str {
		if s == '1' {
			one++
		} else {
			zero++
		}
	}
	return zero, one
}
