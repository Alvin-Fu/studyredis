package main

import (
	"redisstruct"
	"fmt"
	"container/list"
)

func main(){
	l := redisstruct.CreateList()
	l = l.PushBack(1)
	fmt.Println(l)
	l = l.PushBack(2)
	fmt.Println(l)
	l = l.PushBack(2)
	fmt.Println(l)
	fmt.Println(list.New())
}

