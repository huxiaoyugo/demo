package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)


func main() {
	Single()
}

var Address = "39.107.111.1"
var conn net.Conn



func Single() {
	var  err error

	conn, err = net.Dial("tcp", Address + ":8811")  //开启TCP连接端口
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte("data12345678090"))

	time.Sleep(time.Second*10)

	conn.Write([]byte("continue...lllllwww"))

	fmt.Println("成功")
}



func Many() {
	var  err error
	conn, err = net.Dial("tcp", Address + ":8811")  //开启TCP连接端口
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	group := &sync.WaitGroup{}
	for i := 0; i< 4; i++ {
		group.Add(1)
		time.Sleep(time.Second)
		go dialInTCP(i, group)
	}
	group.Wait()
}




func dialInTCP(num int, group *sync.WaitGroup) {
	defer func() {
		group.Done()
	}()

	msg := []byte(fmt.Sprintf("%d: woqunidaye",num))
	if tconn, ok := conn.(*net.TCPConn); ok {   //这里可以断言，因为是tcp连接，其实值类型就是TCPConn
		fmt.Println(num, "assert success")
		//tconn.CloseRead()   //尝试调用TCPConn的函数
		_, err := tconn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
