package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28


问题2：
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。
 */

 func main() {
	 //fmt.Println(uniquePaths2(3,7))
 	fmt.Println(uniquePaths2(5,5))
	 grid := [][]int{

	 }

	 fmt.Println(uniquePathsWithObstacles(grid))
 }

 /*
 直接递归
  */
func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}
	n1 := uniquePaths(m-1,n)
	n2 := uniquePaths(m, n-1)
	return n1+n2
}


func uniquePaths2(m int, n int) int {

	cache := make(map[string]int)
	return uniquePathsFunc(m,n,&cache)
}

func uniquePathsFunc(m, n int, cache *map[string]int) (res int) {

	key1 := fmt.Sprintf("%d-%d", m,n)
	key2 := fmt.Sprintf("%d-%d", n,m)
	if r, ok := (*cache)[key1]; ok {
		return r
	}
	if r, ok := (*cache)[key2]; ok {
		return r
	}

	defer func() {
		(*cache)[key1] = res
		(*cache)[key2] = res
	}()

	if m == 1 || n == 1 {
		res = 1
		return
	}

	n1 := uniquePathsFunc(m-1,n,cache)
	n2 := uniquePathsFunc(m, n-1,cache)
	res = n1 + n2
	return
}



func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cache := make(map[string]int)
	return uniquePathsWithObstaclesFunc(obstacleGrid,0,0,&cache)
}

func uniquePathsWithObstaclesFunc(obstacleGrid [][]int,m, n int, cache *map[string]int)(res int) {
	key1 := fmt.Sprintf("%d-%d", m,n)
	if r, ok := (*cache)[key1]; ok {
		return r
	}

	defer func() {
		(*cache)[key1] = res
	}()

	row := len(obstacleGrid)
	if row == 0 {return 0}
	col := len(obstacleGrid[0])
	if col == 0 {return 0}

	if obstacleGrid[m][n] == 1 {
		res = 0
		return
	}

	if m == row-1 && n == col-1 {
		res = 1
		return
	}

	if m == row-1 {
		return uniquePathsWithObstaclesFunc(obstacleGrid, m,n+1,cache)
	}
	if n == col-1 {
		return uniquePathsWithObstaclesFunc(obstacleGrid, m+1,n,cache)
	}
	n1 := uniquePathsWithObstaclesFunc(obstacleGrid, m+1,n,cache)
	n2 := uniquePathsWithObstaclesFunc(obstacleGrid, m,n+1,cache)
	res = n1 + n2
	return
}