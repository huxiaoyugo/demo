package main

import (
	"fmt"
	"sort"
)

// T[]int 每个小孩的过河时间
/*
首先对T[]进行升序排序
f(1) = T[1]
f(2) = T[2]
f(3) = T[2]+T[1]+T[3]
f(n) = 2*T[2]+T[1]+T[n]+f(n-2) n>=4

 */
func GoRiver(T []int) int {

	sort.Sort(IntArr(T))
	n := len(T)
	if n <= 2 {return T[n-1]}
	if n == 3 {
		return T[1]+T[2]+T[0]
	}
	f_n2 := T[1] // f(2)
	f_n1 := T[1]+T[2]+T[0] // f(3)
	f_n := 0
	for i:=3; i< n; i++ {
		f_n = 2*T[1] + T[0] +T[i] + f_n2
		f_n2 = f_n1
		f_n1 = f_n
	}
	return f_n
}


type IntArr []int
func(s IntArr) Len()int {return len(s)}
func(s IntArr) Swap(i,j int) {s[i],s[j] = s[j],s[i]}
func(s IntArr) Less(i, j int) bool {return s[i]<s[j]}

func main() {
	fmt.Println(GoRiver([]int{1,4,3,2,5}))
}
