package main

import "fmt"


type M struct {
	R,C int
}
func main() {


	var j = 0
	switch j {
	case 0,1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

}


var List =[2]int{}
const (
	A = len(List)+iota //0

	B        // 1
	C  //c
	D = "D"       // c，与上  相同。
	E = iota // 4，显式恢复。注意计数包含了 C、D 两 。
	F        // 5

)

type Po struct {
	S,E int
}

func maxPro(A []int, B []int) []int {

	for i := 1; i < len(A); i++ {
		if A[i] > B[i] {
			// a <= b中寻找
			for k := 0; k < i; k++ {
				if A[k] <= B[k] && A[k] > B[i] && A[i] > B[k] {
					A[k], A[i] = A[i], A[k]
					break
				}
			}
		} else {
			// 在全部中找
			for k := 0; k < i; k++ {
				if A[k] > B[i] && A[i] > B[k] {
					A[k], A[i] = A[i], A[k]
					break
				}
			}

			// 在 a<=b中找
			for k := 0; k < i; k++ {
				if A[k] <= B[k] && (A[k] > B[i] || A[i] > B[k]) {
					A[k], A[i] = A[i], A[k]
					break
				}
			}
		}
	}

	return A
}



