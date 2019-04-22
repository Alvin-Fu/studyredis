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
	iter := redisstruct.CreateIter(l,redisstruct.HEAD_DIRECTION)
	node := redisstruct.MyListNext(iter)
	fmt.Println("node",node.GetValue())
	fmt.Println(iter)
	fmt.Println(l.Len)
}

