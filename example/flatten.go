package example

import (
	. "leetcode/example/nodes" // . 取消前缀
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 * 栈的使用,利用栈的lifo ,实现前序遍历
 *
 */
func flatten(root *TreeNode) {
	list := []*TreeNode{}
	stack := []*TreeNode{}

	node := root
	//先序遍历，把结果放入list。
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		/*
		* 先把后入的节点，拿出后入节点的右边。
		* 在拿右边姐弟拿的左节点。继续 中-左-右
		 */
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[:len(stack)-1]
	}

	/*
	* 原地变成无LeftNode 的二叉树
	 */
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}

}
