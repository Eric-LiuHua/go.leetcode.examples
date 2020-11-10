package tree

import (
	. "leetcode/example/nodes"
)

/**
 * 145. 二叉树的后序遍历
 * 父>右>左。
 */
func postorderTraversal(root *TreeNode) []int {
	sliceNode := make([]*TreeNode, 0)
	res := make([]int, 0)

	if root == nil {
		return res
	}
	sliceNode = append(sliceNode, root)
	for len(sliceNode) > 0 {
		cur := sliceNode[len(sliceNode)-1]
		sliceNode = sliceNode[:len(sliceNode)-1]

		if cur != nil {
			///在右节点之前重新插入该节点，以便在最后处理（访问值）
			//nil跟随t插入，标识已经访问过，还没有被处理
			sliceNode = append(sliceNode, cur, nil)
			if cur.Right != nil {
				sliceNode = append(sliceNode, cur.Right)
			}
			if cur.Left != nil {
				sliceNode = append(sliceNode, cur.Left)
			}

		} else if len(sliceNode) > 0 {
			tmp := sliceNode[len(sliceNode)-1]
			sliceNode = sliceNode[:len(sliceNode)-1]
			res = append(res, int(tmp.Val))
		}

	}
	return res

}
