package main

import (
	"net"
	"time"
)

func handleConnection(c net.Conn) {
	//读写数据
	buffer := make([]byte, 1024)
	c.Read(buffer)
	c.Write([]byte("Hello from server"))
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err!= nil {
			return
		}
		c.Write([]byte("hello"))
		go handleConnection(c)
	}
	time.Sleep(10*time.Second)
}
