package main

import "fmt"

/*

C[i] = {Vi, Wi}  ==> Vi表示第i个物品的价值， Wi表示第i个物品的体积

maxW 表示最大体积

f(c[] , maxW) int ==> 输入C[]和最大体积maxW,得出背包所能够承装的最大价值。
 */

func main() {


	c := []VW{
		{4,3},
		{11,5},
		{8,4},
	}

	fmt.Println(maxVal(c, 17003))


}

type VW struct {
	V int // 价值
	W int // 体积
}


func maxVal(c []VW, maxW int) int {
	r := make([]int, maxW+1)
	for i := 1; i<=maxW; i++ {
		q := getSignalVal(&c, i)
		for j:=1; j<=i/2; j++ {
			q = maxItem(q, r[j] + r[i-j])
		}
		r[i] = q
	}
	//
	//for index, val := range r {
	//	if index == 0 {continue}
	//	fmt.Printf("%d : %d\n", index, val)
	//}
	return r[maxW]
}


// 获取不超过某个重量的最大价值的物品
func getSignalVal(c *[]VW, targetW int) (res int){
	for _, item := range *c {
		if item.W > targetW {
			continue
		}
		if res < item.V {
			res = item.V
		}
	}
	return
}
func maxItem(items ...int) int {

	if len(items) ==0 {
		return 0
	}
	res := items[0]
	for i:=1; i< len(items);i++ {
		if items[i]>res {
			res = items[i]
		}
	}
	return res
}