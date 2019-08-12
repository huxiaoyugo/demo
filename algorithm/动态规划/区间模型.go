package main

import "fmt"

/*
	p[i,j] = 代表的是s[i:j]子串需要插入最少字符数

	if s[i] == s[j] {
		p[i,j] = p[i+1,j-1]
	} else {
		p[i,j] = min(p[i+1,j], p[i,j-1]) + 1
	}

 */

func addCharToHuiWenZi(s string) int {

	count := len(s)
	if count == 1 {
		return 0
	}

	if s[0] == s[count-1] {
		return addCharToHuiWenZi(s[1:count-1])
	} else {
		left := addCharToHuiWenZi(s[0:count-1])
		right := addCharToHuiWenZi(s[1:count])
		if left > right {
			return right+1
		}
		return left + 1
	}
}


func main() {
	fmt.Println(addCharToHuiWenZi("abcaa"))
}
