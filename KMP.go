package main

import "fmt"

func main() {

	//fmt.Println(next("ababa"))

	i := KMP("ababakkkababall","ababa")

	fmt.Println(i)

}
func KMP(str, search string) int {

	next := next(search)
	fmt.Println(next)
	j := 0
	i := 0
	t := len(search)
	for ;i< len(str); {
		if j == -1 || search[j] == str[i] {
			j++
			i++
			if j >= t {
				break
			}
		} else {
			j = next[j]
		}

	}


	if j >= t {
		return i-j
	}
	return -1

}


func next(search string)[]int {

	next := make([]int, len(search))

	next[0] = -1
	for i := 1; i< len(search); i++ {
		sub := search[0:i]
		subLastIndex := len(sub)-1
		for j:=1; j< len(sub); j++ {
			if sub[0:j] == sub[subLastIndex-j+1:] {
				next[i] = j
			}
		}
	}
	return next
}