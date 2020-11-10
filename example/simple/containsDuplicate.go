package simple

//存在重复元素
func containsDuplicate(nums []int) bool {
	var res bool = false
	if len(nums) <= 1 {
		return res
	}
	var m map[int]bool = make(map[int]bool)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			res = true
			break
		} else {
			m[num] = true
		}
	}
	return res
}
