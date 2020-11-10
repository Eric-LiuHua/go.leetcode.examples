package example

import (
	. "leetcode/example/nodes" // . 取消前缀
)

/**
 * 剑指 Offer 55 - I. 二叉树的深度
 * 广度搜索
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var nodes []*TreeNode = make([]*TreeNode, 0)
	var res int = 0
	if root != nil {
		nodes = append(nodes, root)
		for {
			tmpCount := len(nodes)
			for i := tmpCount; i > 0; i-- {
				tmpNode := nodes[0]
				nodes = nodes[1:]
				if tmpNode.Left != nil {
					nodes = append(nodes, tmpNode.Left)
				}
				if tmpNode.Right != nil {
					nodes = append(nodes, tmpNode.Right)
				}

			}
			res++
			if len(nodes) == 0 {
				break
			}
		}
	}
	return res

}
