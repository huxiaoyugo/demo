package CheckTing

import (
	"fmt"
	"testing"
)


func TestGetRest(t *testing.T) {
	test1()
}


func test1() {

	pieces := []int{
		1,1,1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8,
	}
	ting :=CanTing(pieces)
	fmt.Println(ting)
}