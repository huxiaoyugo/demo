package main

import (
	"errors"
	"fmt"
	"runtime"
)

func foo(a, b int) (i int, err error) {
	
	
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err *error) { fmt.Printf("second defer err %v\n", *err) }(&err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}


func defer1() {
	
	a := []int{1}
	defer fmt.Println(a)
	defer func(a []int) {
		fmt.Println(a)
	}(a)

	defer func() {
		fmt.Println(a)
	}()

	 a = append(a, 2,3,4)
	 panic("aaa")
}


func test() {

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("test==%v\n", err)
	//	}
	//}()
	var run func() = nil

	defer func() {
		fmt.Println("runsssss")
	}()
	defer run()
	fmt.Println("runs")
}

func TryC(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}


func test2() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		fmt.Println(recover())
		panic("defer panic")
	}()

	panic("test panic")
}

func test3() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")

		}()
		l:=recover() //无效！
		fmt.Println(l)
	}()

	panic("test panic")
}


func Try(fn func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 1<<16)
			buf = buf[:runtime.Stack(buf, true)]
			switch typ := e.(type) {
			case error:
				err = typ
			case string:
				err = errors.New(typ)
			default:
				err = fmt.Errorf("%v", typ)
			}
			fmt.Printf("==== STACK TRACE BEGIN ====\npanic: %v\n%s\n===== STACK TRACE END =====", err, string(buf))
		}
	}()
	fn()
	return
}


func main() {


	test3()
}