package main

import "fmt"

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
todo: 超时了，需改进
 */

func divide(dividend int, divisor int) int {
	if divisor == dividend {
		return 1
	}
	// 符号
	sign := dividend^divisor > 0
	dividend =abs(dividend)
	divisor = abs(divisor)

	if divisor > dividend  {
		return 0
	}
	res := 0
	t:=divisor
	delt := divisor
	count := 1
	for {
		t+=delt
		res += count
		if t>dividend {
			if count == 1 {break}
			t -= delt
			res -= count
			count = count>>1
			delt = delt>>1
		} else {
			delt = delt<<1
			count = count<<1
		}
	}
	if sign {
		if res > 1<<31-1 {
			return 1<<31-1
		}
	} else {
		if res > 1<<31 {
			return 1<<31-1
		}
		res = -res
	}
	return res
}


func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}


func main() {
	fmt.Println(divide(1<<31,-2))
}