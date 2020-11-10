package dynamic_programming

//300. 最长上升子序列
//给定一个无序的整数数组，找到其中最长上升子序列的长度。不是连续的
func LengthOfLIS(nums []int) int {
	res := 0
	if len(nums) <= 0 {
		return res
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res = 1
	for i := 1; i < len(nums); i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if tmp < dp[j] {
					tmp = dp[j]
				}
			}
		}
		dp[i] = tmp + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
