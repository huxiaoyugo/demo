package CheckTing

import (
	"sort"
	"fmt"
)

var (
	restCount  = 0
	checkCount = 0
)




func CheckTing(pieces []int, ting *[]int) {
	checkCount++
	if len(pieces) == 1 {
		fmt.Println("剩下1个：", pieces)
		*ting = append(*ting, pieces[0])
		return
	}

	restArr, out := GetRest(copySlic(pieces))

	if len(pieces) == 4 {
		fmt.Println("剩下4个：", pieces)
		// 只能是2+2的组合
		flag := false
		if pieces[0] == pieces[1] {
			*ting = append(*ting, getHuP(pieces[2], pieces[3])...)
			flag = true
		}
		if pieces[2] == pieces[3] {
			*ting = append(*ting, getHuP(pieces[0], pieces[1])...)
			flag = true
		}
		if !flag && len(restArr) == 0 {
			fmt.Println("失败")
		}
	}

	for index, arr := range restArr {
		fmt.Println("取出：", out[index])
		fmt.Println("输入：", arr)
		CheckTing(arr, ting)
	}
	return
}

func GetRest(arry []int) ([][]int, [][]int) {
	restCount++
	res := make([][]int, 0)
	out := make([][]int, 0)
	// 从第一个开始找
	for index := 0; index < len(arry); index++ {
		preArr := copySlic(arry[:index])
		arr := copySlic(arry[index:])
		// 判断是否能找到ABC组合
		var count = findFirst(arr, 5)
		// 说明存在ABC
		if count[0] > 0 && count[1] > 0 && count[2] > 0 {
			//fmt.Println("0:ABC")
			out = append(out, []int{arr[0], arr[count[0]], arr[count[0]+count[1]]})

			res = append(res, append(preArr, sliceDel(arr, 0, count[0], count[0]+count[1])...))

			// 判读外层的ABC是否会收到内层ABC的影响
			if count[1] < 2 || count[2] < 2 {
				// 可能会收到影响
				if count[3] > 0 {
					out = append(out, []int{arr[count[0]], arr[count[0]+count[1]], arr[count[0]+count[1]+count[2]]})
					res = append(res, append(preArr, sliceDel(arr, count[0], count[0]+count[1], count[0]+count[1]+count[2])...))
				}
			}
			if count[2] < 2 {
				if count[3] > 0 && count[4] > 0 {
					out = append(out, []int{arr[count[0]+count[1]], arr[count[0]+count[1]+count[2]], arr[count[0]+count[1]+count[2]+count[3]]})
					res = append(res, append(preArr, sliceDel(arr, count[0]+count[1], count[0]+count[1]+count[2], count[0]+count[1]+count[2]+count[3])...))
				}
			}

			if count[1] == 3 {
				t := count[0]
				out = append(out, []int{arr[t], arr[t+1], arr[t+2]})
				res = append(res, append(preArr, sliceDel(arr, t, t+1, t+2)...))
			}
			if count[2] == 3 {
				t := count[0] + count[1]
				out = append(out, []int{arr[t], arr[t+1], arr[t+2]})
				res = append(res, append(preArr, sliceDel(arr, t, t+1, t+2)...))
			}
		}
		if count[0] == 3 { // 检查是否有AAA组合  // 判断是否满足AAA
			out = append(out, []int{arr[0], arr[1], arr[2]})
			res = append(res, append(preArr, sliceDel(arr, 0, 1, 2)...))
		}
		if len(res) > 0 {
			break
		}
	}
	return res, out
}

func findFirst(arr []int, n int) []int {
	var count = []int{1, 0, 0, 0, 0}
	var p = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count[p]++
		} else if arr[i]-1 == arr[i-1] {
			p++
			count[p]++
			if p == n-1 {
				break
			}
		} else {
			break
		}
	}
	return count
}

func getHuP(a, b int) (res []int) {
	if a > b {
		a, b = b, a
	}
	if a == b {
		res = append(res, a)
		return
	}
	if a == b-1 {
		if a-1 > 0 {
			res = append(res, a-1)
		}
		if b+1 < 10 {
			res = append(res, b+1)
		}
		return
	}
	if a == b-2 {
		res = append(res, a+1)
	}
	return
}

func copySlic(arr []int) ([]int) {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func sliceDel(arr []int, indexs ...int) []int {
	slice := copySlic(arr)
	sort.Sort(IntSlice(indexs))
	for _, item := range indexs {
		if item >= len(slice) {
			continue
		}
		if item == len(slice)-1 {
			slice = slice[:item]
		} else {
			slice = append(slice[:item], slice[item+1:]...)
		}
	}
	return slice
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
