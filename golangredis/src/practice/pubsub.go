//redis学习过程中的实践
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func Subs(str string) {  //消费者
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(str) //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("hello %s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string)  { //生产者
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	str := make([]string,0)
	str = append(str,"channel.md", "channel.text", "channel.log")
	_,err1 := conn.Do("PUBLISH", str, message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}

}

func main()  {
	//需要先启动消费者
	go Subs("channel1")
	//go Subs("channel2")
	go Push("hello world h")
	time.Sleep(time.Second*3)
}