package goSdk

import (
	"testing"
	"fmt"
)

func TestFloat64bits(t *testing.T) {

	val := Float64bits(1)
	fmt.Println(val)
}


func TestPointer2Uintptr(t *testing.T) {
	Pointer2Uintptr()
}


func TestReflectPointer(t *testing.T) {
	ReflectPointer()
}

func TestTestCopy(t *testing.T) {
	TestCopy()
}

func TestIllegalUse(t *testing.T) {
	illegalUseA()
	illegalUseB()
}


func TestApply(t *testing.T) {
	Apply2()
}

func TestAlign(t *testing.T) {
	Align()
}


func TestSizeOf(t *testing.T) {
	SizeOf()
}

func TestOffsetof(t *testing.T) {
	Offsetof()
}