package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

func main(){
	//reflect.TypeOf()
	//var a byte
	//a = 1
	//fmt.Println( a&1)
	//fmt.Println(reflect.TypeOf(1<<0).Kind().String())
	//fmt.Println(1<<2)
	var n int64 = 257
	var pn = &n
	var pf = (*int8)(unsafe.Pointer(pn))
	fmt.Println(n)
	// now, pn and pf are pointing at the same memory address
	fmt.Println(*pf) // 2.5e-323
	*pf = 122
	fmt.Println(n) //
	reflect.TypeOf(2)
	reflect.ValueOf(3)
}
