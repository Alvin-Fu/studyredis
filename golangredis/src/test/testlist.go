package main

import (
	"fmt"
	"list"
	"redisstruct"
)

func main(){
	//testPushFront()
	//testPushBack()
	testSPushBack()
	testSPushFront()
}
//测试链表的头插
func testPushFront(){
	l := list.CreateList()
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)
	tmp := l.GetHead()
	for i := 0; i < l.GetLen(); i++{
		fmt.Println(tmp.Value)
		tmp = tmp.NextNode()
	}
}
//测试链表的尾插
func testPushBack(){
	l := list.CreateList()
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	tmp := l.GetHead()
	iter := list.CreateIter(l,1)
	node := iter.Next()
	fmt.Println(node.Value)
	fmt.Println(iter.Next().Value)
	for i := 0; i < l.GetLen(); i++{
		fmt.Println(tmp.Value)
		tmp = tmp.NextNode()
	}
}

//双链表尾插
func testSPushBack(){
	l := slist.CreateList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(3)
	tmp := l.Head
	for i := 0; i < l.Len; i++{
		fmt.Println(tmp.Value,tmp.Next)
		tmp = tmp.GetNext()
	}
	iter := slist.CreateIter(l,2)
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())
}
//双向链表头插
func testSPushFront(){
	l := slist.CreateList()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	tmp := l.Head
	for i := 0; i < l.Len; i++{
		fmt.Println(tmp.Value,tmp.Next)
		tmp = tmp.GetNext()
	}
	iter := slist.CreateIter(l,1)
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())
	fmt.Println("迭代器遍历",iter.MyListNext().GetValue())

}