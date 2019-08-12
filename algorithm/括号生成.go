package main

import (
	"fmt"
)

/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

 func main() {
	arr := generateParenthesis(16)
	//for _, it := range arr {
	//	fmt.Println(it)
	//}
	 fmt.Println(len(arr))
 }

func generateParenthesis(n int) []string {

	res := make([]string, 0)
	gene(&res,"",n,n)
	return res
}

func gene(res *[]string, str string, left, right int) {

	if left == 0 {
		for i := 0; i < right;i++ {
			str += ")"
		}
		*res = append(*res, str)
		return
	}

	gene(res, str+"(", left-1,right)
	if left < right {
		gene(res, str+")", left,right-1)
	}
}