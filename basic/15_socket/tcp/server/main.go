package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP Server端

//处理连接
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		data, err := reader.Read(buf[:]) // 从conn读取数据
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:data]) //转换成string
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed ,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed ,err:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理连接
	}
}
