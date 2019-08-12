package main

import "fmt"

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输入: false
 */


 // todo: 虽然能得到结果，但是时间和空间效率都较低， 需要进一步研究
 func main() {
 	//s := `abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb`
 	s :="aabaabbaaabaaaaaaaaaaaabbabbbbaaaaaaaaaababbbababbbbabbbbbbbbaabaabbaaabaaaaaaaaaaaabbabbbbaaaaaaaaaababbbababbbbabbbbbbbb"
 	//p := "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"

 	fmt.Println(len(s))
 	fmt.Println(isMatch("aa", "*"))
 }


 func isMatch(s string, p string) bool {
	 tMap := make(map[string]bool)
	 return Mactch(s, p, &tMap)
 }


 /*
 子定向下的递归调用，只是使用ma[i-j]记录中间状态， 避免重复计算,
 但是空间和时间复杂度还是较高
  */
func Mactch(s string, p string, tMap *map[string]bool) (res bool) {

	sLen := len(s)
	pLen := len(p)
	if it, ok := (*tMap)[fmt.Sprintf("%d-%d", sLen, pLen)]; ok {
		return it
	}

	defer func() {
		(*tMap)[fmt.Sprintf("%d-%d", sLen, pLen)] = res
	}()

	if pLen == 0 {
		res = sLen == 0
		return
	}
	if sLen == 0 {
		for _, it := range p {
			if it != '*' {
				res = false
				return
			}
		}
		res = true
		return
	}
	var i, j = 0, 0
	// p为字母
	if p[j] >= 'a' && p[j]<='z' {
		if s[i] == p[j] {
			res = Mactch(s[i+1:],p[j+1:],tMap)
			return
		} else {
			res = false
			return
		}
	}

	// p为？
	if p[j] == '?' {
		res=Mactch(s[i+1:],p[j+1:],tMap)
		return
	}

	// p为*
	if p[j] == '*' {
		// 分为三种情况
		// 1、已经是最后一个字母了
		if j == pLen - 1 {
			res = true
			return
		}
		// 判断下一个字母
		switch p[j+1] {
		case '*':
			res = Mactch(s, p[j+1:],tMap)
			return
		case '?':
			for index, _ := range s {
				res = Mactch(s[index:],p[j+1:],tMap)
				if res {
					return
				}
			}
			res = false
			return
		default:
			// 默认为字母
			// *a 需要寻找所有的s中所有的a
			for index, item := range s {
				if uint8(item) == p[j+1] {
					res = Mactch(s[index:],p[j+1:],tMap)
					if res {
						return
					}
				}
			}
			res = false
			return
		}
	}
	res = true
	return
}


