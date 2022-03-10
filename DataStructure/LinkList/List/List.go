package List

type Object interface {}

type Node struct {
	Data Object // 定义数据域
	Next *Node // 定义地址域
}

type List struct {
	headNode *Node // 定义头节点
}

// IsEmpty 判断是否为空链表
func (l *List) isEmpty() bool  {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

// getLinkListLength 获取链表长度
func (l *List) getLinkListLength() int  {
	// 获取头节点
	curNode := l.headNode
	// 定义计数器
	counter := 0

	for curNode != nil {
		counter ++
		curNode = curNode.Next
	}
	return counter
}

func (l *List) add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}