package divideAndConquer

import (
	"testing"
	"fmt"
)

func TestGetOne(t *testing.T) {

	str := "|(&(f,t,t),&(t))"
	//expr := []byte(str)
	//str = string(expr[2:len(expr)-1])

	//fmt.Println(str)
	fmt.Println(parseBoolExpr(str))

	//fmt.Println(getColArr("t,f,t"))
}


func parseBoolExpr(expression string) bool {

	if expression == "t" {return true}
	if expression == "f" {return false}

	expr := []byte(expression)

	if expr[0] == '!' {
		return !parseBoolExpr(string(expr[2:len(expr)-1]))
	}

	if expr[0] == '&' {
		str := string(expr[2:len(expr)-1])

		expArr := getColArr(str)

		for _, item := range expArr {
			if !parseBoolExpr(item) {
				return false
			}
		}
		return true
	}

	if expr[0] == '|' {
		str := string(expr[2:len(expr)-1])
		expArr := getColArr(str)

		for _, item := range expArr {
			if parseBoolExpr(item) {
				return true
			}
		}
		return false
	}

	return false
}


func getColArr(str string) (res []string) {

	leftCount := 0
	col := ""
	for _, item := range []byte(str) {

		if item == ',' && leftCount == 0 {
			res = append(res, col)
			col = ""
		} else {
			col += string(item)
		}

		if item == '(' {
			leftCount += 1
		}

		if item == ')' {
			leftCount -= 1
		}
	}
	res = append(res, col)
	return
}