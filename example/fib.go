package example

import (
	"fmt"
)

//斐波那契数列 ，使用 1000000007 求余 。0 下标是辅助的，所以切片长度要 n+1
//消耗空间换时间
func Fib(n int) int {
	if n < 2 {
		return n
	}

	arr := make([]int, n+1, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % 1000000007
	}
	return arr[n]
}

//使用常量级的赋值。
func FibA(N int) int {
	if N <= 1 {
		return N
	}
	if N == 2 {
		return 1
	}
	var (
		r  int = 0
		p1 int = 1
		p2 int = 1
	)

	for i := 3; i <= N; i++ {
		r = (p1 + p2) % 1000000007
		p1 = p2
		p2 = r
	}
	return r
}
func TestFib() {
	fmt.Printf("TestFib.Fib(100)= %d \n", Fib(100))
	fmt.Printf("TestFib.FibA(100)= %d \n", FibA(100))
}
