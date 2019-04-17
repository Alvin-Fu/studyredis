package redisstruct


//代表一个节点
type MyListNode struct{
	prev, next *MyListNode
	value interface{}
}
func (m *MyListNode) Prev()*MyListNode{
	if p := m.prev; p != nil {
		return p
	}
	return nil
}
func (m *MyListNode) Next()*MyListNode{
	if p := m.next; p != nil {
		return p
	}
	return nil
}

//连表
type MyList struct{
	head, tail *MyListNode		//头结点和尾节点
	len int64					//链表的长度
}
func (l *MyList) PushBack(v interface{})*MyList{
	//if l.tail.next == nil {
		tmp := &MyListNode{
			prev:l.tail,
			next:nil,
			value: v,
		}
		l.tail.next = tmp
		l.len++
	//}
	return l
}


//链表迭代器
type MyListIter struct {
	next *MyListNode	//下一个节点
	index int64 	//方向
}

func CreateList()*MyList{
	myList := new(MyList)
	myList.head = new(MyListNode)
	myList.tail = myList.head
	myList.len = 0
	return myList
}
