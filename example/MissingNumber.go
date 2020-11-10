package example

//剑指 Offer 53 - II. 0～n-1中缺失的数字
func MissingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}
