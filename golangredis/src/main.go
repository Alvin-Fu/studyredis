package main

import (
	"fmt"
	"redisstruct"
)

func main(){
	l := redisstruct.CreateList()
	l = l.PushBack(1)
	fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	l = l.PushBack(2)
	fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	l = l.PushBack(3)
	fmt.Println(l.Head.Value,l.Tail.Value,l.Len)
	fmt.Println(l.Head.Next.Value)
	iter := redisstruct.CreateIter(l,redisstruct.HEAD_DIRECTION)
	node := redisstruct.MyListNext(iter)
	for {
		if node != nil {
			fmt.Println("node next:",node.Value,node)
			//fmt.Println(node.Next.Value)

		} else{
			break
		}
		node = node.Next
	}
	//fmt.Println("node----",node.GetValue())
	fmt.Println(iter)
	fmt.Println(l.Len)
}

