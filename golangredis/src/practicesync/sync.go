package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var mu = new(sync.Mutex)
	//这里面是一个接口，mutex是对这个接口的实现
	cond := sync.NewCond(mu)
	for i := 0; i < 3; i++{
		//wa.Add(1)
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(i)
			//wa.Done()
		}(i)
	}
	time.Sleep(2*time.Second)
	//wa.Wait()
	cond.Signal()
	//唤醒所有等待的goroutine
	cond.Broadcast()
	time.Sleep(2*time.Second)

}



