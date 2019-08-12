package main

import "fmt"

func main() {

	//fmt.Println(expands("abaaba", 2,3))

	fmt.Println(longestPalindrome("AB"))


}

// 中心扩展法
// 以每个字母或者相邻的两个字母为中心，向外搜索

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	var start, end = 0, 0
	for index, _ := range s {
		len1 := expands(s, index, index)
		len2 := expands(s, index, index+1)
		le := len1
		if len2 > le {
			le = len2
		}
		if end-start +1 < le {
			if len1 > len2 {
				start = index - le/2
				end = index + le/2
			} else {
				start = index - le/2 + 1
				end = index + le/2
			}
		}
	}
	return s[start:end+1]
}

func expands(s string, i, j int) (le int) {
	left := i
	right := j
	count := len(s)
	for {
		if left<0 || right>=count || s[left] != s[right] {
			break
		}
		left--
		right++
	}
	left++
	right--
	le = right-left+1
	return
}