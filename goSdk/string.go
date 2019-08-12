package goSdk

import (
	"fmt"
	"unsafe"
)


type stringStruct struct {
	str unsafe.Pointer
	len int
}


type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func String() {
	s1 := "string"

	var bytes []byte
	s2 := s1
	s3 := s1[3:]

	stringS := *(*stringStruct)(unsafe.Pointer(&s1))

	*(*slice)(unsafe.Pointer(&bytes)) = slice{stringS.str, stringS.len, stringS.len}


	fmt.Println(bytes)

	fmt.Println(bytes)

	printAddr := func(sp *string) {
		stringS := (*stringStruct)(unsafe.Pointer(sp))
		fmt.Println(stringS.str, ":", *sp)
	}

	printAddr(&s1)
	printAddr(&s2)
	printAddr(&s3)
}


func String2(str string){
	s1 := str

	var bytes []byte

	stringS := *(*stringStruct)(unsafe.Pointer(&s1))


	*(*slice)(unsafe.Pointer(&bytes)) = slice{stringS.str, stringS.len, stringS.len}
	fmt.Println(bytes)
	bytes[0]='c'
	fmt.Println(bytes)

}


func Copy() {

	var b1 = []byte("abc")
	var b2 = make([]byte,len(b1)+1)

	co := copy(b2, b1)
	fmt.Println(co)

	fmt.Println(b1)
	fmt.Println(b2)


	b2[0] = 'p'

	fmt.Println(b1)
	fmt.Println(b2)

}

func Const() {
	const (
		A   int = 1+iota
		B
		C
		E
		F
		G
		H

		D   = 5
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(E)

	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(D)

}