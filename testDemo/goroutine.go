package main

import (
	"fmt"
	"runtime"
)

var (
	flag = false
	str  string
)

func foo1() {
	fmt.Println("kakkkk")
	flag = true
	str = "setup complete!"

	fmt.Println("kakkkk")
}

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}


func main() {

	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	//go foo1()
	//go foo1()
	//time.Sleep(time.Second*5)

	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(flag)
	for flag {

	}
	//for i := 0; i < 2; i++ {
	//	<- quit
	//}
}
