package dynamic_programming

//答案需要取模 1e9+7（1000000007）
func NumWays(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}
