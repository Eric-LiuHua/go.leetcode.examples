package leetcode

import (
	"fmt"
	"leetcode/example"
	"leetcode/example/nodes"
	"leetcode/example/simple"
	"testing"
)

func Test_main(t *testing.T) {
	fmt.Printf("----------------- leetcodt.main ----------------- \n")

	//example.TestMinimumTotal()
	example.TestFib()

	//要求有序数组
	fmt.Printf("example.TwoSum([]int{2,3, 7, 8, 11, 15},9):%v \n", example.TwoSum([]int{2, 3, 7, 8, 11, 15}, 9))
	fmt.Printf("example.TwoSum([]int{2,7,11,15},9):%v \n", example.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("example.OneTimesNumber([]int{2, 7, 11, 15}):%v \n", example.OneTimesNumber([]int{2, 7, 11, 2, 7, 11, 15}))
	fmt.Printf("example.OneTimesNumber([]int{2, 7, 11, 15}):%v \n", example.OneTimesNumber(nil))

	l1 := nodes.NewNode(2)
	l1.AppendNodes(3, 4, 5, 6, 7, 8, 9)
	l2 := nodes.NewNode(1)
	l2.AppendNodes(1, 1, 11, 12)

	head := example.MergeTwoLists(l1, l2)

	head.Traverse()
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("example.MaxProduct(nums):", example.MaxProduct(nums))
	example.TestLRUCache1()

	fmt.Println("example.YhsjA(3)")
	fmt.Println(example.YhsjA(3))

	fmt.Println("example.YhsjB(3)")
	fmt.Println(example.YhsjB(3))

	numsa := []int{1, 1, 2, 3, 4, 5}
	fmt.Println("simple.RemoveDuplicates(numsa):", simple.RemoveDuplicates(numsa))

	numsb := []int{2, 7, 131, 22, 11, 23}
	fmt.Println("simple.TwoSum(numsb,9):", simple.TwoSum(numsb, 9))

	fmt.Printf("----------------- leetcodt.main ----------------- \n")
}
