package main

import "fmt"

type Index struct {
	I, J int
}

func palindromePartition(s string, k int) int {

	p := geneP(s)

	dp := make([][]int, len(s)+1)
	for i:=0; i< len(dp);i++ {
		dp[i] = make([]int, len(s)+1)
		for j:=0; j< len(dp[i]);j++ {
			dp[i][j] = 9999
		}
	}
	dp[0][0] = 0
	for i:= 0; i< len(s); i++ {
		for j := 1; j<=k; j++ {
			for x := 0; x < i; x++ {
				c := dp[x][j-1] + p[Index{x,i-1}]
				if c < dp[i][j] {
					dp[i][j] = c
				}
			}
		}
	}
	return dp[len(s)][k]
}

func geneP(s string) map[Index]int {
	P := make(map[Index]int,0)
	for i:=0; i<len(s); i++ {
		for j:=i; j<len(s);j++ {
			P[Index{i, j}] = toCycleStrNeedChangeCount(s[i:j+1])
		}
	}
	fmt.Println(P)
	return P
}


func toCycleStrNeedChangeCount(str string) int {
	left := 0
	right := len(str)-1
	count := 0
	for left <= right {
		if str[left] != str[right] {
			count++
		}
		left++
		right--
	}
	return count
}


func main() {
	s := "fyhowoxzyrincxivwarjuwxrwealesxsimsepjdqsstfggjnjhilvrwwytbgsqbpnwjaojfnmiqiqnyzijfmvekgakefjaxryyml"
	k := 32

	fmt.Println( palindromePartition(s,k))
}

