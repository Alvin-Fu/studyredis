package main

import (
	"fmt"
	list "golist"
)

func main(){
	l := list.New()
	fmt.Println("l:",l)
	////获取下一个
	e1 := l.PushBack(2)
	fmt.Println("e1:",e1)
	fmt.Println(e1.GetList())
	e2 := l.PushBack(1)
	fmt.Println("e2",e2)
	fmt.Println(e2.GetList())
	e3 := l.PushFront(2)
	fmt.Println("e3:",e3)
	fmt.Println(e3.GetList())
	e4 := l.PushFront(5)
	fmt.Println(e4.GetList())
	e5 := l.PushFront(5)
	fmt.Println(e5)
	fmt.Println(e5.GetList())
	//fmt.Println(l.Front().Value)
	//fmt.Println(l.Back().Value)

	//var c , d, e = 0, 10, 20
	//Add(c,d,e,add(10,20))
	//fmt.Println(l.Back())
	//fmt.Println(l.Len())

	//fmt.Println(l.Back().Value)
	//fmt.Println(l.Front().Value)
	//l.InsertBefore(6,l.Front())
	//fmt.Println(l.Front().Value)
	//e := l.Remove(l.Front())
	//fmt.Println(l.Front().Value)
	//fmt.Println(e)
	var u UserInfo
	u.GetInfo(10,add)
}


//go中的匿名函数
type Info interface {
	GetInfo(uid int, add func(a, b int)int)
}

type UserInfo struct{
	uid int
	score int
}
func  add(a, b int)int{
	return a + b
}

func (u *UserInfo) GetInfo(uid int, add func(a, b int)int){
	u.uid = uid
	u.score = add(1,2)
	fmt.Println(u)
}
