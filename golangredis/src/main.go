package main

import (
	"redisstruct"
	"fmt"
	"container/list"
)

func main(){
	l := redisstruct.CreateList()
	fmt.Println(l)
	fmt.Println(list.New())
}

