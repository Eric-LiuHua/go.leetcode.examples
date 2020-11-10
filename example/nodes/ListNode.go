package nodes

import "fmt"

//List Node
type ListNode struct {
	Value int64
	Next  *ListNode
}

//初始化
func NewNode(value int64) *ListNode {
	return &ListNode{value, nil}
}

//追加
func (head *ListNode) AppendNode(value int64) *ListNode {
	fmt.Printf("AppendNode :%d \n", value)

	p := head
	tmp := NewNode(value)

	if p == nil {
		head = tmp
	} else {
		for {
			if p.Next == nil {
				break
			}
			p = p.Next
		}
		p.Next = tmp
	}
	return head
}

func (head *ListNode) AppendNodes(value ...int64) *ListNode {
	p := head
	if p == nil && len(value) > 0 {
		head = NewNode(value[0])
		p = head
	} else {
		for {
			if p.Next == nil {
				break
			}
			p = p.Next
		}
	}
	for i := 1; i < len(value); i++ {
		p.Next = NewNode(value[i])
		p = p.Next
	}
	return head
}

func (head *ListNode) Add(value int64) *ListNode {
	p := NewNode(value)
	p.Next = head
	return p
}

//遍历链表
func (head *ListNode) Traverse() {
	p := head
	i := 0
	for {
		if p == nil {
			break
		}
		fmt.Printf("index:%d,value:%d \n", i, p.Value)
		p = p.Next
		i += 1
	}
}

//获取值
func (head *ListNode) GetValue(index int) int64 {
	p := head
	var res int64 = 0
	i := 1
	if p != nil {
		for {
			if p == nil {
				break
			}

			if i == index {
				res = p.Value
			}
			p = p.Next
			i += 1
		}
	}
	return res
}

//删除
func (head *ListNode) Delete(index int) bool {
	p := head
	res := false
	i := 1
	if p != nil {
		for {
			if p == nil {
				break
			}
			//找到上个节点就可以开始处理
			if i == index-1 {
				if p.Next == nil {
					p.Next = nil
				} else {
					p.Next = p.Next.Next
				}
			}

			p = p.Next
			i += 1
		}
	}
	return res
}

//链表反转
func (head *ListNode) Reverse() *ListNode {
	h := head
	var tmp, n *ListNode = nil, nil
	for {
		if h == nil {
			break
		}
		//先临时记录下个节点
		tmp = h.Next
		//先把h的下个节点指向res，在调换
		h.Next = n
		n = h
		//还原节点，继续循环
		h = tmp
	}
	return n
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var cur, res, tmp *ListNode = head, nil, nil
	for {
		if cur == nil {
			break
		}
		//先临时记录下个节点
		tmp = cur.Next
		//先把h的下个节点指向res，在调换
		cur.Next = res
		res = cur
		//还原节点，继续循环
		cur = tmp
	}
	return res
}
func TestMain() {
	//结构，就算是nil，也可以调用结构的方法
	var tmp *ListNode = nil
	tmp = tmp.AppendNodes(3)
	fmt.Println("tmp.Traverse() 1  ************************* ")
	tmp.Traverse()

	ln := NewNode(2)
	ln.AppendNode(3).AppendNode(13).AppendNode(23).AppendNode(33)
	ln.AppendNodes(3, 4, 5, 6, 7, 8, 9)

	fmt.Println("ln.Traverse() 1  ************************* ")
	ln.Traverse()

	fmt.Printf("ln.GetValue(3) :%d  \n", ln.GetValue(3))

	fmt.Printf("ln.Delete(4) :%v  \n", ln.Delete(4))
	fmt.Println("ln.Traverse() 2  ************************* ")
	ln.Traverse()

	ln = ln.Reverse()
	fmt.Println("ln.Traverse() 3  ************************* ")
	ln.Traverse()
}
