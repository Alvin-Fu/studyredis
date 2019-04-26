package main

import "fmt"

type MyTmp interface {
	SetUid(uid int)
	SetTid(tid int)
}


type Tmp struct {
	Uid int
	Tid int
}
func (t *Tmp)SetUid(uid int){
	t.Uid = uid
}
func (t *Tmp)SetTid(tid int){
	t.Tid = tid
}

func main(){
	var tmp = make([]Tmp,0)
	for {
		i := 0
		if len(tmp) > 4{
			break
		}
		t := Tmp{}
		t.SetTid(10+i)
		t.SetUid(10+i)
		tmp = append(tmp,t)
		i++
	}
	Test(tmp)
	fmt.Println(tmp)
}


func Test(tmp []Tmp){
	fmt.Println(tmp)
	var p = make([]Tmp,2)
	//p = tmp
	p = append(p, tmp...)
	p[0].SetUid(8)
	fmt.Println(p)
	fmt.Println(tmp)
}