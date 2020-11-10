package tree

import (
	. "leetcode/example/nodes"
)

/**
 * 144. 二叉树的前序遍历
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
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
			if top.Left != nil {
				sliceNode = append(sliceNode, top.Left)
			}
			sliceNode = append(sliceNode, top, nil)
		} else if len(sliceNode) > 0 {
			tmp := sliceNode[len(sliceNode)-1]
			sliceNode = sliceNode[:len(sliceNode)-1]
			res = append(res, int(tmp.Val))
		}
	}
	return res
}
