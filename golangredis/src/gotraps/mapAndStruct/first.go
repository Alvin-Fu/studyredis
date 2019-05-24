package main

import "fmt"

//关于map中的地址问题

type MyMap struct {
	str string
}

func main(){
	//var m1 = make(map[int]*MyMap)
	//m1[1] = &MyMap{
	//	str: "hello",
	//}
	//m1[1].str = "world"
	//fmt.Println(m1[1].str)
	done := false
	go func() {
		done = true
	}()
	for !done{
		fmt.Println("done")
	}
	fmt.Println("hello")
}




