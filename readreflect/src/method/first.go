package main

import (
	"reflect"
	"fmt"
)

type Read interface{
	MyRead()
}
type read struct{
	r int
}
func (r *read)MyRead(){
	r.r = 10
}

func main(){
	var r read
	fmt.Println(reflect.TypeOf(r).Method(0))

}

