package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 */

 func main() {

 	fmt.Println(lengthOfLongestSubstring("dvdfg"))
 }
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	sByte := []byte(s)
	count := len(sByte)

	if count < 2 {
		return count
	}

	var start, end = 0,1
	for {
		if end >= count {
			break
		}
		for i:= start; i< end; i++ {
			if sByte[i] == sByte[end] {
				if maxLen < end-start {
					maxLen = end-start
				}
				start = i+1
			}
		}
		end++
	}

	// 此处是个需要注意的地方，容易遗漏
	// 就是在循环中没有触发最大值的更新
	// 而是触发了end >= count
	if maxLen < end-start {
		maxLen = end-start
	}
	return maxLen
}