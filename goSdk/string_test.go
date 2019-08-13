package goSdk

import (
	"fmt"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	String()
}

func TestString2(t *testing.T) {
	String2("aabc")
}

func TestConst(t *testing.T) {
	Const()
}

func TestCopy1(t *testing.T) {
	Copy()
}


func TestQuote(t *testing.T) {
	fmt.Println(strconv.QuoteRuneToASCII('c'))
}