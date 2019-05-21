/*死锁
 *
 */
package main

import (
	"fmt"
	"sync"
)
type values struct{
	mu sync.Mutex
	value int
}
func TestMutes(v1, v2 *values){
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()
	//time.Sleep(2*time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()
	fmt.Println("v1+v2=:",v1.value+v2.value)
}

var wg sync.WaitGroup
func main(){
	var a = values{
		value:1,
		mu:sync.Mutex{},
	}
	var b = values{
		value: 2,
		mu:sync.Mutex{},
	}
	wg.Add(2)
	go TestMutes(&a,&b)
	go TestMutes(&b,&a)
	wg.Wait()



	TestDtat()
	//data := 0
	//w := sync.WaitGroup{}
	//for i := 0; i < 100000; i++{
	//	w.Add(1)
	//	go func() {
	//		data++
	//		w.Done()
	//	}()
	//}
	//w.Wait()
	//fmt.Println("data:",data)
}

func TestDtat(){
	var data int
	go func(){
		data++
	}()
	if data == 0{
		fmt.Printf("The data:%d\n",data)
	} else{
		fmt.Printf("the data:%d\n",data)
	}
}
