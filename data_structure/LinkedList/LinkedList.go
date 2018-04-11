package NodeList

// 单个节点
type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

// 从头插入数据
func(list *LinkedList) InsertFirst(v interface{}) {
	// 新建节点
	data := &Node{data: v}

	// 已经存在元素的情况 需要将头指向新的头
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

func(list *LinkedList) InsertLast(v interface{}) {
	// 新建节点
	data := &Node{data: v}
	// 不存在节点的情况直接赋值到head
	if list.head == nil {
		list.head = data
	}
	// 获取最后一个数据
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	// 插入节点
	cur.next = data

}
func(list *LinkedList) RemoveByIdx(idx int) bool {
	// 处理越界情况
	if (list.head == nil) {
		return false
	}
	if (idx < 0) {
		return false
	}
	// 处理情况为 移除首元素的情况
	if (idx == 0) {
		list.head = list.head.next
		return true
	}

	// 处理核心情况
	cur := list.head
	for u := 1; u < idx; u++ {
		cur = cur.next
	}
	//  上一个元素的next指针 = 下一个元素的指针
	cur.next = cur.next.next
	return true	                                                                                                                                                                                                                                                                                                                                                                                                                                                
}
// 插入一个元素
// func (list *LinkedList)AddByIdx(i int, v interface{}) bool {

// }

// 获取前面一个值
func (list *LinkedList)GetFirst(i interface{}) (interface{}, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

// 获取最后一个值
func (list *LinkedList)GetLast(i interface{}) (interface{}, bool) {
	if list.head == nil {
		return 0, false
	}
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur.data, true
}

// 获取长度
func (list *LinkedList)GetSize(i interface{}) int {
	count := 0
	cur := list.head
	for cur != nil {
		count += 1
		cur = cur.next
	}
	return count
}

// 获取值的slice
func (list *LinkedList)GetValues(i interface{}) []interface{} {
	var vals []interface{}
	cur := list.head
	for cur != nil {
		vals = append(vals, cur.data)
		cur = cur.next
	}
	return vals
}
