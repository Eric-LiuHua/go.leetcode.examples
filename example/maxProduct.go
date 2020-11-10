package example

/**
* 1464. 数组中两元素的最大乘积
* 给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 (nums[i]-1)*(nums[j]-1) 取得最大值。
* 请你计算并返回该式的最大值。
*
* 提示：
* 2 <= nums.length <= 500
* 1 <= nums[i] <= 10^3
* 不需要特判，
**/
func MaxProduct(nums []int) int {
	var (
		f, s int = 0, 0
	)
	for _, num := range nums {
		if num >= s {
			s = num
		}
		if s > f {
			s, f = f, s
		}
	}
	return (f - 1) * (s - 1)
}
