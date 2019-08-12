package main

import (
	"time"
	"fmt"
	"sync"
)

func main () {
	timeout := make(chan bool)


	lock := sync.Mutex{}

	go func() {
		time.Sleep(time.Second*20)
		timeout <- true
	}()

	ch := make(chan int)
	over := make(chan bool)

	oh := make(chan string)
   var i  =1

	go func() {
		for ; i<30;  {
			//time.Sleep(time.Second*1)
			//ch <- i

			lock.Lock()
			ch <- i
			i++
			lock.Unlock()
		}
		//over <- true
	}()

	go func() {
		for ; i<30;  {
			//time.Sleep(time.Second*1)


			lock.Lock()
			ch <- i
			i++
			lock.Unlock()
		}
		//over <- true
	}()


   go func() {
   	for  {
		select {
   		case char := <- oh:
   			fmt.Println("jkkf",char)
		}
	}
   }()

	for {
		select {
		case char := <- ch:
			oh <- "hello"
			fmt.Println(char)
			continue
		case <- over:
			fmt.Println("over")
		case <- timeout:
			fmt.Println("time over")
		}
		break
	}

	close(ch)
	close(timeout)
	close(over)
}


func fun() {
	timeout := make(chan bool, 1)
	ch := make(chan string,1)
	over := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second*5)
		timeout <- true
	}()


	go func() {
		for count := 1; count < 3; count ++ {
			time.Sleep(time.Second*1)
			ch <- "huxiaoyu"
		}
		over <- true
	}()

	//for {
		select {
		case <- ch:
			// 从ch中读取到数据
			fmt.Println("数据接收完毕")

		case <- timeout:
			// 没有从ch中读取到数据，但从timeout中读取到了数据
			fmt.Println("time out")

		case <- over:
			fmt.Println("over")

		}
	//}


}

