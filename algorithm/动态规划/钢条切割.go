package main

import (
	"fmt"
)

var (
	p = []int{0,1,5,8,9,10,17,17,20,24,30}
	cacheMap map[int]int
)

func init() {
	cacheMap = make(map[int]int)
}

func main() {
	fmt.Println(maxValue(10))
	fmt.Println(maxValue2(10))
}


/*
分析得到递推公式
f(1) = p(1)
f(2) = max(p(2), f(1)+f(1))
f(3) = max(p(3), f(2)+f(1))
f(n) = max(p(n), f(n-1)+f(1), ... , f(n-n/2)+f(n/2))
 */

// 自定向下的方法
// 任然是递归的方式，只是缓存了中间结果
func maxValue(n int) int  {

	if n <= 1 {
		return n
	}
	if val, ok := cacheMap[n]; ok {
		return val
	}

	m := 0
	if n > 10 {
		for i := 1; i<=10 ;i++ {
			m = max(m, maxValue(n-i) + maxValue(i))
		}
	} else {
		for i := 1; i<n ;i++ {
			m = max(m, maxValue(n-i) + maxValue(i))
		}
		m = max(m, p[n])
	}
	cacheMap[n] = m
	return m
}


// 自底向上的方法
//

func maxValue2(n int) int {
	r := make([]int, len(p))
	for i:=1; i<=n; i++ {
		q := p[i]
		for j:=1; j <= i/2; j++ {
			q = max(q, r[i-j]+r[j])
		}
		r[i] = q
	}
	return r[n]
}


func max(items ...int) int {

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


