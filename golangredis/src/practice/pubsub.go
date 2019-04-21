//redis学习过程中的实践
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	conn, err := redis.Dial("tcp","192.168.75.134:6379")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	pubSub := redis.PubSubConn{
		Conn: conn,
	}
	err = pubSub.Ping("ping")
	if err != nil {
		fmt.Println(err)
	}
}

