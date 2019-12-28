package main

import "fmt"

func main() {
	//p := uint(1)
	//fmt.Println(p | 1<<'a')

	fmt.Println(maxFreq("aababcaab", 2, 3,4))
	//check("abca",2 )
}




func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {

	m := make(map[string]int)
	max := 0
	for i:= 0; i<= len(s); i++ {

		for l:=i+minSize; l<len(s) && l <= i+maxSize;l++ {
			sub := s[i:l]
			if m[sub] > 0 {
				m[sub] += 1
			} else {
				if check(sub, maxLetters) {
					m[sub] = 1
				}
			}
			if m[sub] > max {
				max = m[sub]
			}
		}
	}
	return max
}

func check(s string, maxC int) bool {

	p := uint(0)
	for i:=0; i<len(s); i++ {
		p = p | 1<<(uint(s[i])-97)
	}

	count := 0
	for i:=0;i<28;i++ {
		if p>>uint(i) & uint(1) == 1 {
			count++
		}
	}
	return maxC >= count
}
