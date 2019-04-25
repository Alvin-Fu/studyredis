package redisstruct

<<<<<<< Updated upstream
var (
	HEAD_DIRECTION int64 = 1
	TAIL_DIRECTION int64 = 2
)

=======
import "fmt"
>>>>>>> Stashed changes

//代表一个节点
type MyListNode struct{
	Prev, Next *MyListNode
	Value interface{}
}

func (m *MyListNode) GetValue()interface{}{
	return m.Value
}

func (m *MyListNode) GetPrev()*MyListNode{
	if p := m.Prev; p != nil {
		return p
	}
	return nil
}
func (m *MyListNode) GetNext()*MyListNode{
	if p := m.Next; p != nil {
		return p
	}
	return nil
}

//连表
type MyList struct{
	Head, Tail *MyListNode		//头结点和尾节点
	Len int64					//链表的长度
}



func (l *MyList) PushBack(v interface{})*MyList{
	if l.Tail.Next == nil && l.Len == 0 {
		tmp := &MyListNode{
			Prev:nil,
			Next:nil,
			Value: v,
		}
		l.Head = tmp
		l.Tail = tmp
		l.Len++
	} else {
		tmp := &MyListNode{
<<<<<<< Updated upstream
			Prev:l.Tail,
			Next:nil,
			Value: v,
		}
		t := l.Tail
		l.Tail = tmp
		tmp.Prev = t
		l.Len++
		if l.Len == 2 {
			l.Head.Next = l.Tail
		}
	}
=======
			prev:nil,
			next:nil,
			value: v,
		}
		fmt.Println(tmp)
		tmp.prev = l.tail
		l.head = l.tail
		l.tail.next = tmp
		l.len++
	//}
>>>>>>> Stashed changes
	return l
}


//链表迭代器
type MyListIter struct {
	Next,Prev *MyListNode	//下一个节点
	Direction int64 	//方向
}
func CreateIter(l *MyList,direction int64)*MyListIter{
	myListIter := new(MyListIter)
	if direction == HEAD_DIRECTION{
		myListIter.Next = l.Head
	} else {
		myListIter.Next = l.Tail
	}
	myListIter.Direction = direction
	return myListIter
}
func MyListNext(l *MyListIter)*MyListNode{
	nextNode := new(MyListNode)
	if l.Next != nil {
		if l.Direction == HEAD_DIRECTION {
			nextNode = l.Next
		} else {
			nextNode = l.Prev
		}
	}
	return nextNode
}

func CreateList()*MyList{
	myList := new(MyList)
	myList.Head = new(MyListNode)
	myList.Tail = myList.Head
	myList.Len = 0
	return myList
}
