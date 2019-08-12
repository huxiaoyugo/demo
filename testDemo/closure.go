package main

//
//type Obj struct {
//	Num int
//}

func main() {
	//fmt.Println(*addr1(1))
}

//func getAddFun(obj *Obj)func()int {
//	return func() int {
//		obj.Num++
//		return obj.Num
//	}
//}


//func test1() func(){
//
//	obj := Obj{}
//	num := 0
//	f := func(){
//		obj.Num++
//		num++
//		fmt.Println(num)
//	}
//
//	num+=100
//	f()
//	f()
//	f()
//
//	fmt.Println("obj.Num:", obj.Num)
//	fmt.Println("num:", num)
//
//	return f
//}
//
//func test2() {
//	s := 0
//	for i := 0; i < 3; i++ {
//		j := i
//		s += func(ii int) int {return 2*ii + j }(i)
//	}
//
//	fmt.Println(s)
//}

func addr1(x int) func()int {

	return func() int {
		x++
		return x
	}
}
