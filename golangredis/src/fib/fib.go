package main

import (
	"fmt"
	"time"
)


func main(){
	var n int64 = 30
	t1 := time.Now()
	fmt.Println( Fibo3(n))
	fmt.Println(time.Since(t1))
	t2 := time.Now()
	fmt.Println(Fibo2(n))
	fmt.Println(time.Since(t2))
	t3 := time.Now()
	//Fib(n,fib)
	fib(n)
	fmt.Println(time.Since(t3))

}
//go中的匿名函数，作为参数
func Fib(n int64, fib func(n int64)int64){
	fmt.Println(fib(n))
}

func fib(n int64)int64{
	if n == 0 {
		return 0
	} else if n == 1{
		return 1
	} else if n > 1{
		return fib(n-1) +fib(n-2)
	} else {
		return 0
	}
}
func fib1()func() int64{
	var a, b int64 = 0, 1
	return func() int64 {
		a, b = b, a + b
		return a
	}
}

func Fibo2(n int64) int64 {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		var a, b int64= 1, 1
		var result int64 = 0
		for i := 3; i <= int(n); i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

//匿名函数版
func Fibo3(n int64) int64 {
	if n < 0 {
		return -1
	} else {
		f := Fibonacci()
		var result int64 = 0
		for i := 0; i < int(n); i++ {
			result = f()
		}
		return result
	}
}
//作为返回值
func Fibonacci() func() int64 {
	var a, b int64 = 0, 1
	return func() int64 {
		a, b = b, a+b
		return a
	}
}
