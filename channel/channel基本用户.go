package main

import (
	"fmt"
	"time"
)

func main() {
	t3()
}


/*

知识点：
1、使用空结构体channel传递消息
2、close(chan), 可关闭channel


 */

func t1() {
	done := make(chan struct{}) //发送空结构体（通知）
	c := make(chan string)      //数据传输通道
	go func() {
		s := <-c //接收消息
		println(s)
		done<- struct{}{}
		time.Sleep(time.Second)
		close(done) //关闭通道，为结束通知
		//done<- struct{}{}
	}()

	go func() {
		select {
		case <-done:
			fmt.Println("完成1")
		}
	}()

	go func() {
		select {
		case <-done:
			fmt.Println("完成2")
		}
	}()

	c <- "hi!" //发送消息


	time.Sleep(time.Second*5)
	fmt.Println("结束")
}


func t2(){
	c := make(chan int, 3)		//创建带有3个缓冲区的异步通道
	c <- 1
	c <- 2						//缓冲区没满
	fmt.Println(<-c)				//缓冲区有数据不会阻塞
	fmt.Println(<-c)
	c <- 3
	x, ok := <-c
	if ok { //判断通道是否关闭
		fmt.Println(x)
	}

	for x := range c { //循环获取消息
		fmt.Println(x)
	}
}

//在程序中异步通道可以提高程序的性能减少排队阻塞
//channel变量本身为指针



func t3() {
	c:= make(chan int, 2)
	d:= make(chan int, 5)
	out : for {
		select {
		case c <- 1:
			fmt.Println("add 1")
		case d <- 2:
			fmt.Println("add 2")
		case <-time.After(time.Second):
			fmt.Println("超时")
			break out
		default:
			fmt.Println("default")
			break out
		}
	}
}