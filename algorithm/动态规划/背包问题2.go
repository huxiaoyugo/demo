package main

import "fmt"

func main() {


	wArr := []int{1,2,3,4,3,5,1,3,4,5,6,7,8}
	vArr := []int{1,3,5,9,4,5,7,3,7,8,9,0,1}

	maxVal, res := MaxVal(wArr, vArr, 20)

	fmt.Println("最大价值：", maxVal)
	fmt.Println("选择的物品：", res)
}


func MaxVal(wArr, vArr []int, maxW int)(int, []int) {

	count := len(wArr)
	P := make([][]int, count+1)
	for i:=0; i<=count; i++ {
		P[i] = make([]int, maxW+1)
	}

	for k:=1; k<=count; k++ {
		for w :=1; w<= maxW; w++ {
			v1 := P[k-1][w]
			v2 := 0
			restW := w - wArr[k-1]
			if  restW >= 0 {
				v2 = P[k-1][restW] + vArr[k-1]
			}
			if v2 > v1 {
				P[k][w] = v2
			} else {
				P[k][w] = v1
			}
		}
	}

	for _, row := range P {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println("")
	}

	res := make([]int, 0)
	for k, w := count, maxW;k>0; k--  {
		if P[k][w] != P[k-1][w] {
			res = append(res, k)
			w -= wArr[k-1]
		}
	}
	return P[count][maxW], res
}

