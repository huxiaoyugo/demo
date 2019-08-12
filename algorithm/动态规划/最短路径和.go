package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

 func main() {

 	grid := [][]int{
 		{1,2},
		{3,1},
	}
	fmt.Println(minPathSum(grid))
 }


func minPathSum(grid [][]int) (res int) {

	cache := make(map[string]int)
	return minPathSumFun(grid, 0,0, &cache)
}

func minPathSumFun(grid [][]int, m, n int, cache *map[string]int) (res int) {

	key := fmt.Sprintf("%d-%d", m, n)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	defer func() {
		(*cache)[key] = res
	}()

	row := len(grid)
	col := len(grid[0])

	if m == row - 1 {
		for i := n; i< col; i++ {
			res += grid[m][i]
		}
		return
	}
	if n == col - 1 {
		for i := m; i< row; i++ {
			res += grid[i][n]
		}
		return
	}

	n1 := minPathSumFun(grid,m+1, n, cache)
	n2 := minPathSumFun(grid,m, n+1, cache)

	if n1 > n2 {
		res = grid[m][n] + n2
	} else {
		res = grid[m][n] + n1
	}
	return
}
