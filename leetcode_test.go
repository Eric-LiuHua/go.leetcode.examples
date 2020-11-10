package leetcode

import (
	"leetcode/dynamic_programming"
	"testing"
)

func TestNumWays(t *testing.T) {
	tmp := dynamic_programming.NumWays(92)
	if tmp == 720754435 {
		t.Log(tmp)
	} else {
		t.Fail()
	}
}

func TestLengthOFLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	tmp := dynamic_programming.LengthOfLIS(nums)
	t.Log(tmp)
}
