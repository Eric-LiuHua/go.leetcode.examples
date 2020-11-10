package example

//双向链表
type DListNode struct {
	Key, Value int
	Next, Pre  *DListNode
}

/**
 * 面试题 16.25. LRU缓存
 * 设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。
 * 缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。
 * 当缓存被填满时，它应该删除最近最少使用的项目。
 *
 * 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。
 * 当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 *
 **/
type LRUCache struct {
	m        map[int]*DListNode
	capacity int
	head     *DListNode
}

//初始化构造函数
func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*DListNode), capacity, nil}
}

//新增head
func (this *LRUCache) AddNode(node *DListNode) {
	if this.head == nil {
		this.head = node
		this.head.Next = node
		this.head.Pre = node
	}
}

//更新head
func (this *LRUCache) UpdateHead(node *DListNode, isNew bool) {
	if node != this.head {
		if !isNew {
			//不是新的节点，那就需要把当前节点的pre节点消除。
			node.Pre.Next = node.Next
			node.Next.Pre = node.Pre
		}
		//把当前节点在双链表消除。
		node.Pre = this.head.Pre
		node.Next = this.head
		//把当前节点介入链表，在更换head
		this.head.Pre.Next = node
		this.head.Pre = node
		this.head = node
	}
}

//更新map的缓存数据
func (this *LRUCache) ReplaceNode(key int, value int) {
	if len(this.m) >= this.capacity {
		//消除上个节点的存储。
		delete(this.m, this.head.Pre.Key)
		//把上节点的位置移动到head ，在更新其中的值。
		this.head = this.head.Pre
		this.head.Key = key
		this.head.Value = value
		//新的head 存储到
		this.m[this.head.Key] = this.head
	}
}

//获取数据
func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.UpdateHead(node, false)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.Value = value
		this.UpdateHead(node, false)

	} else {
		if len(this.m) >= this.capacity {
			this.ReplaceNode(key, value)
		} else {
			tmpNode := &DListNode{Key: key, Value: value, Pre: nil, Next: nil}
			if this.head == nil {
				this.AddNode(tmpNode)
			} else {
				this.UpdateHead(tmpNode, true)
			}
			this.m[key] = tmpNode
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
//[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]

func TestLRUCache() {
	var lc LRUCache = Constructor(10)
	lc.Put(10, 13)
	lc.Put(3, 17)
	lc.Put(6, 11)
	lc.Put(10, 5)
	lc.Put(9, 10)
	lc.Get(13)
	lc.Put(2, 19)
	lc.Get(2)
	lc.Get(3)
	lc.Put(5, 25)
	lc.Get(8)
	lc.Put(9, 22)
	lc.Put(5, 5)
	lc.Put(1, 30)
	lc.Get(11)
	lc.Put(9, 12)
	lc.Get(7)
	lc.Get(5)
	lc.Get(8)
	lc.Get(9)
	lc.Put(4, 30)
	lc.Put(9, 3)
	lc.Get(9)
	lc.Get(10)
	lc.Get(10)
	lc.Put(6, 14)
	lc.Put(3, 1)
	lc.Get(3)
	lc.Put(10, 11)
	lc.Get(8)
	lc.Put(2, 14)
	lc.Get(1)
	lc.Get(5)
	lc.Get(4)
	lc.Put(11, 4)
	lc.Put(12, 24)
	lc.Put(5, 18)
	lc.Get(13)
	lc.Put(7, 23)
	lc.Get(8)
	lc.Get(12)
	lc.Put(3, 27)
	lc.Put(2, 12)
	lc.Get(5)
	lc.Put(2, 9)
	lc.Put(13, 4)
	lc.Put(8, 18)
	lc.Put(1, 7)
	lc.Get(6)
	lc.Put(9, 29)
	lc.Put(8, 21)
	lc.Get(5)
	lc.Put(6, 30)
	lc.Put(1, 12)
	lc.Get(10)
	lc.Put(4, 15)
	lc.Put(7, 22)
	lc.Put(11, 26)
	lc.Put(8, 17)
	lc.Put(9, 29)
	lc.Get(5)
	lc.Put(3, 4)
	lc.Put(11, 30)
	lc.Get(12)
	lc.Put(4, 29)
	lc.Get(3)
	lc.Get(9)
	lc.Get(6)
	lc.Put(3, 4)
	lc.Get(1)
	lc.Get(10)
	lc.Put(3, 29)
	lc.Put(10, 28)
	lc.Put(1, 20)
	lc.Put(11, 13)
	lc.Get(3)
	lc.Put(3, 12)
	lc.Put(3, 8)
	lc.Put(10, 9)
	lc.Put(3, 26)
	lc.Get(8)
	lc.Get(7)
	lc.Get(5)
	lc.Put(13, 17)
	lc.Put(2, 27)
	lc.Put(11, 15)
	lc.Get(12)
	lc.Put(9, 19)
	lc.Put(2, 15)
	lc.Put(3, 16)
	lc.Get(1)
	lc.Put(12, 17)
	lc.Put(9, 1)
	lc.Put(6, 19)
	lc.Get(4)
	lc.Get(5)
	lc.Get(5)
	lc.Put(8, 1)
	lc.Put(11, 7)
	lc.Put(5, 2)
	lc.Put(9, 28)
	lc.Get(1)
	lc.Put(2, 2)
	lc.Put(7, 4)
	lc.Put(4, 22)
	lc.Put(7, 24)
	lc.Put(9, 26)
	lc.Put(13, 28)
	lc.Put(11, 26)

}

func TestLRUCache1() {
	var lc LRUCache = Constructor(2)
	lc.Put(1, 1)
	lc.Put(2, 2)
	lc.Get(1)
	lc.Put(3, 3)
	lc.Get(2)
	lc.Put(4, 4)
	lc.Get(1)
	lc.Get(3)
	lc.Get(4)

}
