package main

import "fmt"


func main(){
	var n int64 = 100
	fmt.Println( Fibo3(n))
	fmt.Println(fib(n))
	//Fib(n,fib)

}
//go中的匿名函数
func Fib(n int64, fib func(n int64)int64){
	t := fib(n)
	fmt.Println(t)
}

func fib(n int64)int64{
	var rue int64
	if n <= 1{
		rue = 1
	} else {
		rue = fib(n-1) + fib(n-2)
	}
	return rue
}

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

func Fibonacci() func() int64 {
	var a, b int64 = 0, 1
	return func() int64 {
		a, b = b, a+b
		return a
	}
}
