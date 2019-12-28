package main

import (
	"fmt"
	"sort"
)

func main() {
	routes := [][]int{{1, 2, 3}, {3, 6, 7}}
	S := 1
	T := 6
	fmt.Println(numBusesToDestination(routes, S, T))
}



func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
    p := creatMap(routes, S, T)
	// 广度优先搜索
	return scaleSearch(p)
}

func scaleSearch(p [][]int) int {

	h := make([]int, len(p))
	// 标记已经访问过的节点
	h[len(p)-2] = 1
	stack := make([]int, 0)
	for i, n := range p[len(p)-2] {
		if n == 1 {
			stack = append(stack, i)
		}
	}

	dep := 0
	lastIndex := len(p)-1
	for len(stack) > 0 {
		next := []int{}
		for _, n := range stack {
			if n == lastIndex {
				return dep
			}
			if h[n] != 1 {
				h[n] = 1
				for j, k := range p[n] {
					if k == 1 && h[j] != 1 {
						next = append(next, j)
					}
				}
			}
		}
		dep++
		stack = next
	}
	return -1
}

func creatMap(routes [][]int, S int, T int) [][]int {
	p := make([][]int,len(routes)+2)
	for i:=0; i<len(routes)+2;i++ {
		p[i] = make([]int, len(routes)+2)
	}
	sIndex := len(routes)
	tIndex := len(routes)+1


	for _, l := range routes {
		sort.Ints(l)
	}

	for i, r := range routes {
		if search(r, S) {
			p[i][sIndex] = 1
			p[sIndex][i] = 1
		}
		if search(r, T) {
			p[i][tIndex] = 1
			p[tIndex][i] = 1
		}
	}

	for i:=0; i<len(routes);i++ {
		for j:=i+1; j<len(routes);j++ {
			r1 := routes[i]
			r2 := routes[j]

			// 判断是否有交集
			if inter(r1, r2) {
				p[i][j] = 1
				p[j][i] = 1
			}
		}
	}

	fmt.Println(p)
	return p
}


func inter(r1, r2 []int) bool {
	var i, j int
	for i < len(r1) && j < len(r2) {
		if r1[i] == r2[j] {
			return true
		}
		if r1[i] < r2[j] {
			i++
		} else {
			j++
		}
	}
	return false
}


func search(route []int, val int) bool {
	l := 0
	r := len(route) - 1
	for l <= r {

		mid := (l + r) / 2
		if route[mid] == val {
			return true
		}
		if route[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

