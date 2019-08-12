package main

import (
	"fmt"
	"sort"
)

var (
	F = make([][]int, 0)
)

func main() {

	str := "(((()))((())))"


	//fmt.Println(maxDeepByStack(str))
	//
	//fmt.Println(maxDeepByRecursive(str))

	fmt.Println(maxDepthAfterSplit(str))
}

func maxDepthAfterSplit(seq string) []int {

	path := maxDeepPath(seq)


	fmt.Println(path)
	res := make([]int, len(seq))

	count := len(path) / 2
	if count % 2 != 0 {
		count++
	}

	for i:=0; i< count; i++ {
		res[path[i]] = 1
	}
	return res
}



// 栈
func maxDeepByStack(str string) int {

	stack := 0
	maxDeep := 0
	for _, item := range []byte(str) {
		if item == '(' {
			stack++
			if maxDeep < stack {
				maxDeep = stack
			}
		}
		if item == ')' {
			stack--
		}
	}
	return maxDeep
}

// 递归实现
func maxDeepByRecursive(str string) int {

	if str == "" {
		return 0
	}

	indexs := make([]int, 0)
	stack := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack++
		}
		if str[i] == ')' {
			stack--
		}
		if stack == 0 {
			indexs = append(indexs, i)
		}
	}

	last := 0
	maxD := 0
	for _, itemEnd := range indexs {
		var dep = 0
		if itemEnd - last == 1 {
			dep = 1
		} else {
			dep = 1 + maxDeepByRecursive(str[last+1:itemEnd])
		}
		last = itemEnd + 1
		if dep > maxD {
			maxD = dep
		}
	}
	return maxD
}


// 寻找最深路径
// 使用栈，一次遍历，找到每一个括号对的深度， 并保存在数组P中
// 对P按照前括号编号由小到大进行排序
// 递归找出路径


type IndexRange struct {
	i int
	j int
	Val int
}

func maxDeepPath(str string)[]int {

	var p = make([]IndexRange, 0)
	stack := make([]int,0)
	dep :=0
	for i, item := range str {
		if item == '(' {
			dep = 0
			// 入栈
			stack = append(stack, i)
		} else {
			dep++
			p = append(p, IndexRange{stack[len(stack)-1], i, dep})
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}


	sort.Slice(p, func(i, j int) bool {
		return p[i].i<p[j].i
	})

	for _, item := range p {
		fmt.Printf("(%d,%d)=%d  ", item.i, item.j, item.Val)
	}
	fmt.Println("")

	path := make([]int, 0)

	start := 0
	end := len(p)-1

	for {
		if start >= len(p) || start>end {
			break
		}

		maxStart := start
		maxEnd := start
		for i:=start; i<=end; i++ {

			if p[i].Val > p[maxStart].Val {
				maxStart = i
			}

			if p[i].j < p[maxStart].j {
				maxEnd = i
			}
		}

		path = append(path, p[maxStart].i, p[maxStart].j)
		start = maxStart+1
		end = maxEnd
	}


	return path
}