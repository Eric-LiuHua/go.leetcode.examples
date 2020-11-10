package example

//167. 两数之和 II - 输入有序数组
//二分查找
func TwoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right && numbers[i] <= target {
			mid := ((right - left) >> 1) + left
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < target-numbers[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return []int{-1. - 1}
}
