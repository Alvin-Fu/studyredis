package main

import (
	"fmt"
)

type Receiver struct{
	res string
}
type printer interface {
	First()
}

func (r *Receiver) First(){
	fmt.Println("hello ",r.res)
}
//关于使用指针作为方法的接收者
func main(){
	r := Receiver{
		res: "world",
	}
	r.First()
	var r2 printer = &Receiver{"China"}
	r2.First()
	//会报错，不能取r3的地址，主要是map中的值是不可寻址的
	var r3 = map[string]Receiver{
		"h":Receiver{"BeiJing"},
	}
	fmt.Println(r3)
	//r3["h"].First()
	//不能更新map里面结构体的值
	//r3["h"].res = "shangHai"
}
