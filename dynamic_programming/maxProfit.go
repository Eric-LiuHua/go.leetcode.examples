package dynamic_programming

//只有一个最大值
func maxProfit(prices []int) int {
	maxres := 0
	if len(prices) <= 1 {
		return maxres
	}
	minprice := prices[0]

	//找最小值，在取最大差值作为返回
	for i := 0; i < len(prices); i++ {
		if minprice > prices[i] {
			minprice = prices[i]
		}
		if prices[i]-minprice > maxres {
			maxres = prices[i] - minprice
		}
	}
	return maxres
}

//剑指 Offer 63. 股票的最大利润
func maxProfitOffer(prices []int) int {
	var res int = 0
	if len(prices) > 1 {
		var min int = prices[0]
		for _, v := range prices {
			if v < min {
				min = v
			} else if (v - min) > res {
				res = (v - min)
			}
		}
	}
	return res
}

//122. 买卖股票的最佳时机 II
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//相当于每天都进行交易。
func maxProfitII(prices []int) int {
	var res int = 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}
