package example

import (
	. "leetcode/example/nodes" // . 取消前缀
)

//合并两个有序链表
//需要一个哨兵节点，在结束的时候快速返回链表头。
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Value > l2.Value {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}
