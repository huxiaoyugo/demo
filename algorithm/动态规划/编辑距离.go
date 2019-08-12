package main

import "fmt"

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

 func main() {
 	word1 := "aab"
 	word2 := "abbd"
	 fmt.Println(minDistance(word1, word2))
 }

func minDistance(word1 string, word2 string) int {
	count1 := len(word1)
	count2 := len(word2)
	if count1 == 0 {
		return count2
	}
	if count2==0 {
		return count1
	}
	p := make([][]int, count1+1)
	for i:=0; i<count1+1; i++ {
		p[i] = make([]int, count2+1)
	}
	// 初始化第一行
	for i:=0; i<count2+1; i++ {
		p[0][i] = i
	}
	for i:= 0; i< count1+1; i++ {
		p[i][0]=i
	}

	for i:=1; i<= count1; i++ {
		for j:=1; j<=count2; j++ {
			p1 := p[i-1][j-1]
			p2 := p[i-1][j]+1
			p3 := p[i][j-1]+1
			if word1[i-1] != word2[j-1] {
				p1 +=1
			}
			min := p1
			if p2 < min {
				min = p2
			}
			if p3 < min {
				min = p3
			}
			p[i][j] = min
		}
	}
	return p[count1][count2]
}


