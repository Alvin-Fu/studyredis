package main

import (
	"fmt"
	)

func main(){
	//l := list.New()
	////获取下一个
	//l.PushBack(2)
	//fmt.Println(l.Front().Value)
	//fmt.Println(l.Back().Value)
	var u UserInfo
	u.GetInfo(10,add)
	//var c , d, e = 0, 10, 20
	//Add(c,d,e,add(10,20))
	//fmt.Println(l.Back())
	//fmt.Println(l.Len())
	//l.PushBack(1)
	//l.PushFront(2)
	//l.PushFront(5)
	//fmt.Println(l.Back().Value)
	//fmt.Println(l.Front().Value)
	//l.InsertBefore(6,l.Front())
	//fmt.Println(l.Front().Value)
	//e := l.Remove(l.Front())
	//fmt.Println(l.Front().Value)
	//fmt.Println(e)
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
