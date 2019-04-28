package main

import (
	"list"
	"fmt"
)

func main(){
	myList := linklist.CreateList()
	myList.PushBack(10)
	myList.PushBack(100)
	myList.PushBack(100)
	tmp := myList.Head
	for i := 0; i < myList.Len; i++{
		fmt.Println(tmp.Value)
		tmp = tmp.Next
	}
	//l := redisstruct.CreateList()
	//l = l.PushBack(1)

	//iter := redisstruct.CreateIter(l,redisstruct.HEAD_DIRECTION)
	//fmt.Println(l)
	//fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	//l = l.PushBack(2)
	//fmt.Println(l)
	//fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	//l = l.PushBack(3)
	//fmt.Println(l)
	//fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	//fmt.Println(l.Head.Next.Value)
	//l = l.PushBack(4)
	//fmt.Println(l)
	//fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	//fmt.Println(l.Head.Next.Value)
	//
	//node := redisstruct.MyListNext(iter)
	//fmt.Println("node next:",node.Value,node,&node)

	//for {
	//	if node != nil {
	//		fmt.Println("node next:",node.Value,node)
	//		//fmt.Println(node.Next.Value)
	//
	//	} else{
	//		break
	//	}
	//	node = node.Next
	//}

	//fmt.Println("node----",node)
	//fmt.Println(iter)
	//fmt.Println(l.Len)
	//ln := list.New()
	//fmt.Println(ln)
	//ln.PushBack(1)
	//fmt.Println(ln)
	//ln.PushBack(2)
	//fmt.Println(ln)
	//ln.PushBack(2)
	//fmt.Println(ln)
	//fmt.Println(ln.Front().GetList())
	//fmt.Println(ln.Front().Value,"---",ln.Front().Prev(),"---",ln.Front().Next(),"---",ln.Front().Next().Value,
	//	"---",ln.Front().Next().Prev(),"--",ln.Front().Next().Next().GetList())

}

