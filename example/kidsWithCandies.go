package example

/**
* 1431. 拥有最多糖果的孩子
* 给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
* 对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
 */
func kidsWithCandies(candies []int, extraCandies int) []bool {
	clen := len(candies)
	maxCandies := candies[0]
	//寻找最大值
	for i := 1; i < clen; i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	//通过最大值判断额外糖果加上现有的是否超过最大值，可以多个最大值。所以取等
	var res []bool = make([]bool, clen)
	for i := 0; i < clen; i++ {
		res[i] = (candies[i] + extraCandies) >= maxCandies
	}
	return res
}
