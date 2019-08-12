package main

import "fmt"

func main() {
	var a = 1
	fmt.Println(a)
}


func f() int {
	var res int
	defer func() {
		res++
	}()
	return res
}
