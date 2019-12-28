package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

func groupThePeople(groupSizes []int) [][]int {

	var res = make([][]int, 0)
	var tmpMap = make(map[int]*[]int, 0)

	for index, item := range groupSizes {
		list := tmpMap[item]
		if list == nil {
			list = &[]int{} //&make([]int, 0)
			tmpMap[item] = list
		}
		*list = append(*list, index)
	}

	for k, v := range tmpMap {
		l := make([]int, 0, k)
		for _, item := range *v {
			l = append(l, item)
			if len(l) == k {
				res = append(res, l)
				l = make([]int,0, k)
			}
		}
	}

	return res
}
func findMax(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func find(nums []int, threshold int) int {

	left := 1
	right := findMax(nums)
	for left <= right {
		mid := (left + right) /2
		sum := 0
		for _, num := range nums {
			sum += upDivide(num, mid)
		}

		if sum <= threshold {
			right = mid -1
		}
		if sum > threshold {
			left = mid + 1
		}
	}
	return left
}

func upDivide(a, b int) int {
	if a % b == 0 {
		return a/b
	}
	return a/b + 1
}
func main() {

	nums := []int{2,3,5,7,11}
	old := 5
	fmt.Println(find(nums, old))
	a := []int{1,2,4,3}
	sort.Sort(sort.Reverse(a))

}

var chMap = sync.Map{} // 全局的map,保存请求的chan
func Loop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 获取参数
	uid, _ := strconv.Atoi(r.FormValue("uid"))
	//lastId, _     := strconv.Atoi(r.FormValue("lastId"))
	// 删除消费过的数据,删除<lastId的数据, 检查redis中是否有历史数据,如果有的话,直接返回
	// 如果没有的话
	ch := make(chan []byte)
	chMap.Store(uid, ch)
	select {
	case val := <-ch:
		fmt.Println(val)
		// 写返回的数据等操作
		w.Write(val)
	case <-time.After(30 * time.Second):
		fmt.Println("超时")
	}
	// 删除ch
	chMap.Delete(uid)
	return
}
func sub() {
	// redis.sub("key")  订阅相关消息, 比如支付完成，从订阅中获得的数据，读取uid
	msg := "{uid:10000006,orderId:1009204029,msgId:1023,type:101}"
	uid := 10000006
	if ch, ok := chMap.Load(uid); ok {
		// 发送消息
		ch.(chan []byte) <- []byte(msg)
	}
	// 如果没有连接，不做处理
	return
}
