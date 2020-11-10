package tree

import (
	. "go.leetcode.examples/example/nodes"
)

type TreeNodeStack []*TreeNode

func (tStack *TreeNodeStack) Push(tr *TreeNode) {
	*tStack = append(*tStack, tr)
}
func (tStack *TreeNodeStack) Pop() *TreeNode {
	res := []*TreeNode(*tStack)[len(*tStack)-1]
	*tStack = []*TreeNode(*tStack)[:len(*tStack)-1]
	return res
}
