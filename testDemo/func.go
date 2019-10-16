package main

import (
	"fmt"
)


type I interface {
	testT()
	testP()
}

type S struct {
	*T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}


func methodCollectTest() {
	var a I
	t1 := S{ &T{1}}
	t2 := &t1

	t3 := T{1}
	a = &t3

	a = t1
	a.testT()
	a.testP()
	a = t2
	a.testT()
	a.testP()
}
func main() {


	var i interface{}

	fmt.Println(i)

	var a *T
	i = a

	if i == nil {
		fmt.Println()
	} else {

	}

	if _, ok := i.(int); ok {
		fmt.Println("可以")
	} else {
		fmt.Println("不可以")
	}


}