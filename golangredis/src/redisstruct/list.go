package slist

var (
	HEAD_DIRECTION int = 1
	TAIL_DIRECTION int = 2
)

//实现一个双链表

//代表一个节点
type MyListNode struct {
	Next, Prev *MyListNode
	Value      interface{}
}

func (m *MyListNode) GetValue() interface{} {
	return m.Value
}

func (m *MyListNode) GetPrev() *MyListNode {
	if p := m.Prev; p != nil {
		return p
	}
	return nil
}
func (m *MyListNode) GetNext() *MyListNode {
	if p := m.Next; p != nil {
		return p
	}
	return nil
}

//连表
type MyList struct {
	Head, Tail *MyListNode //头结点和尾节点
	Len        int         //链表的长度
}

func (l *MyList) PushBack(v interface{}) {
	tmp := new(MyListNode)
	tmp.Value = v
	tmp.Prev = nil
	tmp.Next = tmp.Prev
	if l.Len == 0 {
		l.Head = tmp
		l.Tail = tmp
	} else {
		l.Tail.Next = tmp
		tmp.Prev = l.Tail
		l.Tail = tmp
	}
	l.Len++
}
func (l *MyList) PushFront(v interface{}){
	tmp := new(MyListNode)
	tmp.Value = v
	tmp.Prev = nil
	tmp.Next = tmp.Prev
	if l.Len == 0 {
		l.Head = tmp
		l.Tail = tmp
	} else {
		tmp.Next = l.Head
		l.Head.Prev = tmp
		l.Head = tmp
	}
	l.Len++
}


//链表迭代器
type MyListIter struct {
	Next      *MyListNode //下一个节点
	Direction int       //方向
}

func CreateIter(l *MyList, direction int) *MyListIter {
	myListIter := new(MyListIter)
	if direction == HEAD_DIRECTION {
		myListIter.Next = l.Head
	} else {
		myListIter.Next = l.Tail
	}
	myListIter.Direction = direction
	return myListIter
}
func (l *MyListIter) MyListNext() *MyListNode {
	nextNode := new(MyListNode)
	nextNode = l.Next
	if l.Direction == HEAD_DIRECTION {
		l.Next = l.Next.GetNext()
	} else {
		l.Next = l.Next.GetPrev()
	}

	return nextNode
}

func CreateList() *MyList {
	myList := new(MyList)
	tmp := new(MyListNode)
	myList.Head = tmp
	myList.Len = 0
	return myList
}
