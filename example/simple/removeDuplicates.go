package simple

//传引用，删除重复元素，并返回剩下的数量
func RemoveDuplicates(nums []int) int {
	var res int = 0
	var m map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp, ok := m[res]
		if !ok || tmp != nums[i] {
			res = res + 1
			m[res] = nums[i]
		}
	}
	for i := 0; i < res; i++ {
		nums[i] = m[i+1]
	}
	return res
}

//双指针方式
func RemoveDuplicatesA(nums []int) int {
	var res, height int = 0, 1

	for height < len(nums) {
		if nums[res] != nums[height] {
			res++
			nums[res] = nums[height]
		}
		height++
	}

	return res + 1
}
