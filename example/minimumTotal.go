package example

import (
	"fmt"
	"math"
)

//120. 三角形最小路径和
//动态规划，
//先初始化，三角形顶的数据，每下一层。就找该坐标的上一层，前一个位置和当前位置的最小值。加上当前位置的数。
//走完最后一行，及全部统计出路径和。在选择最小的即可
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	fmt.Printf("f  :%v \n", f)

	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, f[n-1][i])
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//测试数据准备
func TestMinimumTotalData() [][]int {
	r := make([][]int, 0)

	c1 := make([]int, 0)
	c1 = append(c1, 2)

	c2 := make([]int, 0)
	c2 = append(c2, 4)
	c2 = append(c2, 3)

	c3 := make([]int, 0)
	c3 = append(c3, 6)
	c3 = append(c3, 5)
	c3 = append(c3, 7)

	c4 := make([]int, 0)
	c4 = append(c4, 4)
	c4 = append(c4, 1)
	c4 = append(c4, 8)
	c4 = append(c4, 3)

	r = append(r, c1)
	r = append(r, c2)
	r = append(r, c3)
	r = append(r, c4)

	return r
}

func TestMinimumTotal() {
	tmp := TestMinimumTotalData()
	fmt.Printf("tmp :=TestMinimumTotalData():%v \n", tmp)

	fmt.Printf("minimumTotal:%d \n", minimumTotal(tmp))

}
