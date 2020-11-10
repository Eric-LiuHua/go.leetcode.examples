package simple

//189. 旋转数组
//借助数组
func rotate(nums []int, k int) {
	var tmp []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
}

//反转，注意编辑，反转过的数，在下次不参与
func rotateA(nums []int, k int) {
	////原始数组                  : 1 2 3 4 5 6 7
	tmp := k % len(nums)
	//反转所有数字后             : 7 6 5 4 3 2 1
	reverse(nums, 0, len(nums)-1)
	//反转前 k 个数字后          : 5 6 7 4 3 2 1
	reverse(nums, 0, tmp-1)
	//反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
	reverse(nums, tmp, len(nums)-1)
}

//反转
func reverse(nums []int, start, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
