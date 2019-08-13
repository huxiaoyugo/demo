package goSdk

import (
	"fmt"
	"sync"
)

func WaitGroup() {


	waitG := sync.WaitGroup{}

	waitG.Wait()
	waitG.Add(1)

	fmt.Println("over")
}
