package simple

//计算数组交集
func intersect(nums1 []int, nums2 []int) []int {

	var dic map[int]int = make(map[int]int)
	var res []int = make([]int, 0)

	for _, v := range nums1 {
		_, ok := dic[v]
		if ok {
			dic[v]++
		} else {
			dic[v] = 1
		}
	}
	for _, v := range nums2 {
		tmp, ok := dic[v]
		if ok && tmp >= 1 {
			res = append(res, v)
			dic[v]--
		}
	}

	return res

}
