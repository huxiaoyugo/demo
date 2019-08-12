package main

import (
	"fmt"
)

func main() {

}

func Lsc(str1, str2 string) int {

	len1 := len(str1)
	len2 := len(str2)

	P := make([][]int, len1+1)
	S := make([][]string, len1+1)
	for i:=0; i<=len1; i++ {
		P[i] = make([]int, len2+1)
		S[i] = make([]string, len2+1)
	}

	for i :=0; i<len1; i++ {
		for j := 0; j<len2; j++ {
			if str1[i] == str2[j] {
				P[i+1][j+1] = P[i][j] + 1
				S[i+1][j+1] = "左上"
			} else {

				P[i+1][j+1] = P[i+1][j]
				S[i+1][j+1] = "左"
				if P[i][j+1] > P[i+1][j] {
					P[i+1][j+1] = P[i][j+1]
					S[i+1][j+1] = "上"
				}
			}
		}
	}

	for i :=1; i<=len1; i++ {
		for j := 1; j<=len2; j++ {
			fmt.Printf("%3d", P[i][j])
		}
		fmt.Println("")
	}


	for i :=1; i<=len1; i++ {
		for j := 1; j<=len2; j++ {
			fmt.Printf("  %s", S[i][j])
		}
		fmt.Println("")
	}

	for i, j := len1, len2; i !=0 && j!=0; {
		switch S[i][j] {
		case "左上":
			fmt.Printf(" %c<-", str1[i-1])
			i -= 1
			j -= 1
		case "上":
			i -= 1
		case "左":
			j -= 1
		}
	}
	return 0

}