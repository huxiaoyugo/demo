package divideAndConquer

import (
	"math/rand"
	"sort"
)

// n个芯片中好的比坏的多，如何快速的找出一片好的芯片
// 1-好 0-坏



func GetOne(cards []int, res *[]int) {


	if len(cards) == 0 {
		return
	}

	var goodIndex = make([]int,0)
	var badCount = make([]int,0)

	for i:= 1; i< len(cards); i++ {
		if test(cards[i], cards[0]) {
			goodIndex = append(goodIndex, i)
		} else {
			badCount = append(badCount, i)
		}
	}

	var needDel []int
	// 说明card[0]是好的
	// 删除说坏的的芯片
	// 并且将这个好的添加到结果集中
	if len(goodIndex) >= len(badCount) {
		*res = append(*res, cards[0])
		// 把自己也添加到要删除的列表中去
		badCount = append(badCount, 0)
		needDel = badCount
	} else {
		// 说明card[0]是坏的
		// 删除说好的的芯片
		// 同时删除这个本生
		goodIndex = append(goodIndex, 0)
		needDel = goodIndex
	}

	GetOne(sliceDel(cards, needDel...), res)
}



func copySlic(arr []int) ([]int) {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func sliceDel(arr []int, indexs ...int) []int {
	slice := copySlic(arr)
	sort.Ints(indexs)
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

// 用a测试b
func test(a, b int) bool {
	if a > 0 {
		return b>0
	}
	return rand.Int() % 2 == 0
}
