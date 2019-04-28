
package list
//迭代器方向
//实现了一个单链表
var (
	HEAD_DIRECTION int = 1
	TAIL_DIRECTION int = 2
)

//代表一个节点
type MyListNode struct{
	NextN *MyListNode
	Value interface{}
}
func (myL *MyListNode) NextNode()*MyListNode{
	if myL.NextN != nil {
		return myL.NextN
	}
	return nil
}
type ListIter struct{
	Iter *MyListNode
	Direction int
}
//创建迭代器
func CreateIter(l *List, direction int)*ListIter{
	iter := new(ListIter)
	if direction == HEAD_DIRECTION{
		iter.Iter = l.head
	}
	iter.Direction = direction
	return iter
}
//获取下一个元素
func (i *ListIter) Next()*MyListNode{
	node := new(MyListNode)
	node = i.Iter
	if i.Direction == HEAD_DIRECTION{
		i.Iter = i.Iter.NextN
	}
	return node
}


type List struct{

	head *MyListNode
	len int
}

func CreateList()*List{
	lt := new(List)
	lt.head = new(MyListNode)
	lt.len = 0
	return lt
}
func (l *List)PushFront(v interface{})*MyListNode{
	if l == nil {
		l = CreateList()
	}
	tmp := new(MyListNode)
	tmp.Value = v
	tmp.NextN = nil
	if l.len == 0{
		l.head = tmp
		l.head.NextN = nil
	} else {
		tmp.NextN = l.head
		l.head = tmp
	}
	l.len++
	return tmp
}
func (l *List) PushBack(v interface{})*MyListNode{
	if l == nil {
		l = CreateList()
	}
	tmp := new(MyListNode)
	tmp.Value = v
	tmp.NextN = nil
	if l.len == 0{
		l.head = tmp
		l.head.NextN = nil
	} else {
		t := l.head
		for i := 0; i < l.len; i ++{
			if l.head.NextN == nil {
				l.head.NextN = tmp
				break
			}
			l.head = l.head.NextN
		}
		l.head = t
	}
	l.len++
	return tmp
}
func (l *List) GetLen()int{
	return l.len
}
func (l *List) GetHead()*MyListNode{
	return l.head
}