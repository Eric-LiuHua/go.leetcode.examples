package tree

import (
	. "go.leetcode.examples/example/nodes"
)

// 二叉树的遍历主要有三种：
// （1）先(根)序遍历（根左右）
// （2）中(根)序遍历（左根右）
// （3）后(根)序遍历（左右根）
/**
 * 94. 二叉树的中序遍历
 * 给定一个二叉树，返回它的中序 遍历。
 * 压栈顺序，右>根>左。
 * 出栈顺序，左>根>右
 */
func inorderTraversal(root *TreeNode) []int {
	sliceNode := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}

	sliceNode = append(sliceNode, root)

	for len(sliceNode) > 0 {
		//顶层节点
		top := sliceNode[len(sliceNode)-1]
		//出栈
		sliceNode = sliceNode[:len(sliceNode)-1]
		// 如果栈的顶层节点不为空，则入栈右左中的顺序
		// 最后一个入栈的同时压入nil标识位
		// 当top为空时，则取出top的下一个节点
		if top != nil {
			if top.Right != nil {
				sliceNode = append(sliceNode, top.Right)
			}
			sliceNode = append(sliceNode, top, nil)
			if top.Left != nil {
				sliceNode = append(sliceNode, top.Left)
			}

		} else if len(sliceNode) > 0 {
			tmp := sliceNode[len(sliceNode)-1]
			sliceNode = sliceNode[:len(sliceNode)-1]
			res = append(res, int(tmp.Val))
		}
	}
	return res
}
