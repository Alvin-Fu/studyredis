package redisstruct

import "fmt"

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
	if l.tail.next == nil && l.len == 0 {
		tmp := &MyListNode{
			prev:nil,
			next:nil,
			value: v,
		}
		l.head = tmp
		l.tail = tmp
		fmt.Println(l)
		l.len++
	} else {
		tmp := &MyListNode{
			prev:l.tail,
			next:nil,
			value: v,
		}
		t := l.tail
		l.tail = tmp
		tmp.prev = t
		l.len++
	}
	return l
}


//链表迭代器
type MyListIter struct {
	next *MyListNode	//下一个节点
	direction int64 	//方向
}
func CreateIter(l *MyList,direction int64)*MyListIter{
	myListIter := new(MyListIter)
	if direction == 1{
		myListIter.next = l.head
	} else {
		myListIter.next = l.tail
	}
	myListIter.direction = direction
	return myListIter
}


func CreateList()*MyList{
	myList := new(MyList)
	myList.head = new(MyListNode)
	myList.tail = myList.head
	myList.len = 0
	return myList
}
