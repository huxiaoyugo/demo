package CheckTing

import (
	"sort"
)

func CanTing(pieces []int) []int {
	sort.Ints(pieces)
	res := make([]int, 0)
	for _, item := range GetMaybeTing(pieces) {
		p := InserToOrderdSlice(pieces, item)
		if CheckHu(p) {
			res = append(res, item)
		}
	}
	return res
}

func InserToOrderdSlice(slice []int, ele int)[]int {

	for index, item := range slice {
		if ele <= item {
			res := append([]int{}, slice[:index]...)
			res = append(res, ele)
			res = append(res, slice[index:]...)
			return res
		}
	}
	return []int{}
}


func CheckHu(pieces []int) bool {

	// 找出对子的位置
	pairIndexs := GetPairIndexs(pieces)

	if len(pairIndexs) == 7 {
		return true
	}

	lastPairVal := -1
	for _, item := range pairIndexs {
		if lastPairVal == pieces[item] {
			continue
		} else {
			// 删除对子
			pList := sliceDel(pieces, item, item+1)
			if IsAllSequenceOrTriplet(pList) {
				return true
			}
		}
	}
	return false
}

func IsAllSequenceOrTriplet(pieces []int) bool {

	count := len(pieces)/3
	for i := 0; i < count; i++ {
		find := FindTripletAndRemove(&pieces)
		if !find {
			find = FindSequenceAndRemove(&pieces)
		}
		if !find {
			return false
		}
	}
	return len(pieces) == 0
}

func FindSequenceAndRemove(piece *[]int) bool {
	resIndex := []int{0}
	for i := 1; i < len(*piece); i++ {
		if (*piece)[i] == (*piece)[i-1] {
			continue
		}
		if (*piece)[i] == (*piece)[i-1]+1 {
			resIndex = append(resIndex, i)
			if len(resIndex) == 3 {
				break
			}
			continue
		}
		return false
	}
	if len(resIndex) != 3 {
		return false
	}
	// 删除
	for i := 2; i >= 0; i-- {
		p := resIndex[i]
		if i == len(*piece)-1 {
			*piece = (*piece)[:p]
		} else {
			*piece = append((*piece)[:p], (*piece)[p+1:]...)
		}
	}
	return true
}

func FindTripletAndRemove(piece *[]int) bool {
	if len(*piece) < 3 {
		return false
	}
	has := IsTriplet((*piece)[0], (*piece)[1], (*piece)[2])
	if has {
		*piece = (*piece)[3:]
	}
	return has
}

func IsTriplet(A, B, C int) bool {
	return A == B && B == C
}

func GetPairIndexs(pieces []int) []int {

	res := make([]int, 0)
	for i:=1; i< len(pieces); i++ {
		if pieces[i] == pieces[i-1] {
			res = append(res, i-1)
			i++
		}
	}
	return res
}

func GetMaybeTing(arr []int) []int {

	res := make([]int, 0)
	last := -2
	for _, item := range arr {
		if last == item {
			continue
		}
		if last+1 == item {
			res = append1to9(res, item+1)
		} else if last+2 == item {
			res = append1to9(res, item, item+1)
		} else {
			res = append1to9(res, item-1, item, item+1)
		}
		last = item
	}
	return res
}

func append1to9(res []int, eles...int) []int{
	for _, item := range eles {
		if item >= 1 && item<=9 {
			res = append(res, item)
		}
	}
	return res
}

func DelSameEle(arr []int) ([]int) {

	slice := make([]int, len(arr))
	copy(slice, arr)
	if len(slice) <= 1 {
		return slice
	}
	sort.Ints(slice)
	p :=1
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			slice[p] = slice[i]
			p++
		}
	}
	slice = slice[:p]
	return slice
}
