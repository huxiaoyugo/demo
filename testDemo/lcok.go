package main

import (
	"bytes"
	"fmt"
	"io"
)

func main(){
	//声明
	//var mutex sync.Mutex
	//fmt.Println("Lock the lock. (G0)")
	////加锁mutex
	//mutex.Lock()
	//
	//fmt.Println("The lock is locked.(G0)")
	//for i := 1; i < 4; i++ {
	//	go func(i int) {
	//		fmt.Printf("Lock the lock.(G%d)\r\n", i)
	//		mutex.Lock()
	//		fmt.Printf("The lock is locked.(G%d)\r\n", i)
	//		mutex.Unlock()
	//	}(i)
	//}
	//time.Sleep(time.Second)
	//fmt.Println("Unlock the lock. (G0)")
	////解锁mutex
	//mutex.Unlock()
	//
	//fmt.Println("The lock is unlocked. (G0)")
	//time.Sleep(time.Second)


	const debug = false
	var buf bytes.Buffer //原来是var buf *bytes.Buffer
	//if debug{
	//	buf = new(bytes.Buffer)
	//}
	f1(&buf)

}
func f1(out io.Writer){
	if out != nil{
		fmt.Println("surprise!")
	}
}