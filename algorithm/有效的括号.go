package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 2:

输入: "()[]{}"
输出: true

示例 4:

输入: "([)]"
输出: false
 */

func isValid(s string) bool {

	stack := make([]int32, 0)

	for _, item := range s {

		switch item {
		case '[', '(','{':
			stack = append(stack, item)
		case ']':
			if len(stack) == 0 {return false}
			it := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			if it != '[' {
				return false
			}
		case ')':
			if len(stack) == 0 {return false}
			it := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			if it != '(' {
				return false
			}
		case '}':
			if len(stack) == 0 {return false}
			it := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			if it != '{' {
				return false
			}
		default:
			return false
		}
	}
	return  len(stack) == 0
}
