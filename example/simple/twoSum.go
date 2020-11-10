package simple

//两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
func TwoSum(nums []int, target int) []int {
	//用于存差额及下标。
	var tmp map[int]int = make(map[int]int)
	var res []int = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		//计算差额
		diff := target - nums[i]
		//找差额在之前是否出现过。出现就返回结果。否则继续循环
		f, ok := tmp[diff]
		if ok {
			res = append(res, f, i)
			break
		} else {
			tmp[nums[i]] = i
		}
	}

	return res
}
