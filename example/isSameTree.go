package example

import (
	. "leetcode/example/nodes"
)

//使用切片实现队列
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var check []*TreeNode = make([]*TreeNode, 0)
	var res bool = (p == nil && q == nil)
	if !res {
		if p == nil || q == nil {
			res = false
		} else {
			if p.Val == q.Val {
				check = append(check, p, q)
				res = true
			}
			for len(check)%2 == 0 && len(check) > 0 && res {
				tlen := len(check)
				for i := 0; i < tlen; i += 2 {
					n0 := check[0]
					n1 := check[1]
					check = check[2:]
					if n1 != nil && n0 != nil && n0.Val == n1.Val {
						check = append(check, n0.Left, n1.Left, n0.Right, n1.Right)
					} else {
						if n1 == nil && n0 == nil {
							res = true
						} else {
							res = false
							break
						}
					}

				}
			}
		}
	}
	return res
}
