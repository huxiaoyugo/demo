package main

import (
	"fmt"
	"reflect"
)


type ManKind interface {
	Run() bool
}

type Man struct {

}

func(m* Man) Run()  bool {
fmt.Println("man run")
return false
}

type Animal struct {
	Name string
}


func (animal *Animal) Quack() string {
	return animal.Name
}

type Duck interface {
	Quack() string

	GetName() string
}

type Cat struct {
	Name string

}


func(c *Cat)setName(name string) {
	c.Name = name
}

func (c Cat)GetName()string {
	return c.Name
}

type Dog struct {
	Cat
}

type Test struct {
	value int
}

func (t *Test) print() {
	println(t.value)
}



func main() {

	fmt.Println(testDefer())

}

func testDefer()(ret int) {

	cat := Cat{}

	defer func() {
		fmt.Println(cat.GetName())
	}()

	cat.Name = "hu"
	return 0
}


func NilOrNot(v interface{}) {


	if reflect.ValueOf(v).IsNil() {
		println("nil")
	} else {
		println("non-nil")
	}
}

func test(i interface{}) {


	fmt.Println(reflect.TypeOf(i))

	//c :=i.(Man)

	switch v:=i.(type) {
	case Cat:
		fmt.Println("cat")
	case *Cat:
		fmt.Println("*cat")
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case int64:
		fmt.Println("int64")

	case struct{}:
		fmt.Println("struct")

	case func():
		fmt.Println("func()")

	default:
		fmt.Println(v)
		fmt.Println("default")
	}
}
