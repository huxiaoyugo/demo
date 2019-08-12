package goSdk

import (
	"unsafe"
	"fmt"
	"reflect"
	"bytes"
	"encoding/gob"
	"github.com/qianlnk/log"
	"strings"
)


func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

type B struct {
	Name string
	Age int
	Score int
}
func Offsetof() {
	b := &B{}
	fmt.Println(unsafe.Offsetof(b.Name)) // 0
	fmt.Println(unsafe.Offsetof(b.Age))  // 16
	fmt.Println(unsafe.Offsetof(b.Score)) // 24
}

func Pointer2Uintptr() {

	var s = "aba"
	o := unsafe.Pointer(&s)

	p :=  unsafe.Pointer(uintptr(o))

	fmt.Println(p == o)

	fmt.Println(p)

	fmt.Println(o)

	end := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s)+1)

	fmt.Println(end)

}

func ReflectPointer() {
	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

	fmt.Println(*p)
}


type A struct {
	Name string
	Age int
}

type I interface {
	GetName() string
	SetName(name string)
}

func(a A) GetName() string{
	return a.Name
}

func(a A) SetName(name string){
	a.Name = name
}

func TestCopy() {

	a := &[]*A{
		&A{Name:"hu",Age:12},
		&A{Name:"ll",Age:13},
	}


	b, err := deepCopy(a)
	if err != nil {
		log.Error(err)
	}


	fmt.Println(a)
	fmt.Println(b)

}



// case A: conversions between unsafe.Pointer and uintptr
//         don't appear in the same expression
func illegalUseA() {
	fmt.Println("===================== illegalUseA")

	pa := new([4]int)

	// split the legal use
	// p1 := unsafe.Pointer(uintptr(unsafe.Pointer(pa)) + unsafe.Sizeof(pa[0]))
	// into two expressions (illegal use):
	//ptr := uintptr(unsafe.Pointer(pa))
	p1 := unsafe.Pointer(uintptr(unsafe.Pointer(pa)) + unsafe.Sizeof(pa[0])*2)
	// "go vet" will make a warning for the above line:
	// possible misuse of unsafe.Pointer

	// the unsafe package docs, https://golang.org/pkg/unsafe/#Pointer,
	// thinks above splitting is illegal.
	// but the current Go compiler and runtime (1.7.3) can't detect
	// this illegal use.
	// however, to make your program run well for later Go versions,
	// it is best to comply with the unsafe package docs.

	*(*int)(p1) = 123
	fmt.Println("*(*int)(p1)  :", *(*int)(p1)) //

	fmt.Println(pa)
}

// case B: pointers are pointing at unknown addresses
func illegalUseB() {
	fmt.Println("===================== illegalUseB")

	a := [4]int{0, 1, 2, 3}
	p := unsafe.Pointer(&a)
	p = unsafe.Pointer(uintptr(p) + uintptr(len(a)-1) * unsafe.Sizeof(a[0]))
	// now p is pointing at the end of the memory occupied by value a.
	// up to now, although p is invalid, it is no problem.
	// but it is illegal if we modify the value pointed by p
	*(*int)(p) = 123
	fmt.Println("*(*int)(p)  :", *(*int)(p)) // 123 or not 123
	// the current Go compiler/runtime (1.7.3) and "go vet"
	// will not detect the illegal use here.

	// however, the current Go runtime (1.7.3) will
	// detect the illegal use and panic for the below code.
	p = unsafe.Pointer(&a)
	for i := 0; i < len(a); i++ {
		*(*int)(p) = 123 // Go runtime (1.7.3) never panic here in the tests

		fmt.Println(i, ":", *(*int)(p))
		// panic at the above line for the last iteration, when i==4.
		// runtime error: invalid memory address or nil pointer dereference

		p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(a[0]))
	}
}



type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}


// 应用：修改其他包中的私有变量
// 方法1：定义一个与其他包中一样的结构体，利用unsafe.Pointer转换成自己定义的结构体指针，然后直接修改
func Apply1() {

	str := "abcdefghijklmn"
	reader := strings.NewReader(str)

	uReader := (*Reader)(unsafe.Pointer(reader))

	by, _ := reader.ReadByte()
	fmt.Println(by)

	uReader.i = 0
	rune, size, _ := reader.ReadRune()
	fmt.Println(rune, size)
}


// 方法2: 使用reflect中字段的StructField 中的offset + 结构体对象的指针值
func Apply2() {
	sr := strings.NewReader("abcdef")
	// 但是我们可以通过 unsafe 来进行修改
	// 先将其转换为通用指针
	// 确定要修改的字段（这里不能用 unsafe.Offsetof 获取偏移量，因为是私有字段）
	if sf, ok := reflect.TypeOf(*sr).FieldByName("i"); ok {
		// 偏移到指定字段的地址
		up := uintptr(unsafe.Pointer(sr)) + sf.Offset
		// 转换为通用指针
		p := unsafe.Pointer(up)
		// 转换为相应类型的指针
		pi := (*int64)(p)
		// 对指针所指向的内容进行修改
		*pi = 3 // 修改索引
	}
	// 看看修改结果
	fmt.Println(sr)
	// 看看读出的是什么
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}


func Align() {
	a := struct {
		I8  int8
		I82 int8
		Name string
		I16 int32

	}{}
	fmt.Println(unsafe.Alignof(a))
	fmt.Println(unsafe.Alignof(a.I8))
	fmt.Println(unsafe.Alignof(a.I16))
	fmt.Println(unsafe.Alignof(a.Name))

	////输出长度为16
	//fmt.Println(unsafe.Sizeof(struct {
	//	p   *int8
	//	i8  int8
	//}{}))
}


func deepCopy(src interface{}) (interface{}, error) {

	typ := reflect.TypeOf(src)

	isPtr := false
	switch typ.Kind() {
	case reflect.Ptr:
		typ = typ.Elem()
		isPtr = true
		break
	}

	fmt.Println(typ)
	dst := reflect.New(typ).Interface()
	fmt.Println(reflect.TypeOf(dst))


	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return nil, err
	}

	err := gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
	if err !=nil {
		return nil, err
	}

	if !isPtr {
		dst = reflect.ValueOf(dst).Elem().Interface()
	}

	return dst, nil
}


func SizeOf() {

	var a int32
	var b = `abcdefghijklmnopqfdsfdsfdsafdsfdsafdsfdsfdsfdsfds`
	var c = [10]int{1,2,3}
	var d = []int{1,2,3,4,5}
	var f = 'a'

	fmt.Println(unsafe.Sizeof(a)) // 4
	fmt.Println(unsafe.Sizeof(b)) // 16
	fmt.Println(unsafe.Sizeof(c)) // 80
	fmt.Println(unsafe.Sizeof(d)) // 24
	fmt.Println(unsafe.Sizeof(f)) // 4

	/*
	type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
	}
	 */
}