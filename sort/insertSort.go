package main

import "fmt"

func main() {

	s := []int{1,4,5,2,3,8,4,5}
	InsertSort2(s)
	fmt.Println(s)


	c := []int{1,2,2,2,2,3}
	fmt.Println(findIndex(c, len(c)-1, 2))

}


func InsertSort(s []int) {

	if len(s) <= 1 {
		return
	}
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j+1] < s[j] {
				s[j+1],s[j] = s[j],s[j+1]
			} else {
				break
			}
		}
	}
}

func InsertSort2(s []int) {

	if len(s) <= 1 {
		return
	}
	for i := 1; i < len(s); i++ {
		j := findIndex(s, i-1, s[i])
		t := s[i]
		for k := i - 1; k >= j; k-- {
			s[k+1], s[k] = s[k], s[k+1]
		}
		s[j] = t
	}
}


func findIndex(s []int, end int, val int) int {

	start := 0
	for ;start<=end; {
		mid := (start + end) /  2
		if s[mid] > val {
			end = mid -1
		} else {
			start = mid + 1
		}
	}
	return start
}