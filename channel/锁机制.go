package main

import (
	"sync"
	"fmt"
	"strings"
)

func main() {

	//Mutex()
	//Channel()

	str :=`[["azdwg"],["ohmvnkxnokpnwxmr","mljb"],["jdjdb"],[],[],[],["xvn","mljb"],[],["xvn"],[],["xvn"],["jdjdb","mljb"],[],["ohmvnkxnokpnwxmr","mljb","ztp"],["ohmvnkxnokpnwxmr"],[],[],["ohmvnkxnokpnwxmr","jdjdb"],["jdjdb"],[],[],[],["azdwg","mljb"],["ztp"],["xvn"],["jdjdb"],["mljb"],["xvn"],["rdaavykiextpr"],[],[],["ohmvnkxnokpnwxmr","jdjdb","mljb","rdaavykiextpr"],[],["jdjdb","mljb","rdaavykiextpr","ztp"],["ohmvnkxnokpnwxmr"],["rdaavykiextpr","ztp"],["mljb"],["rdaavykiextpr"],[],["azdwg"],[],["rdaavykiextpr"],["jdjdb"],[],["mljb","ztp"],[],["rdaavykiextpr","ztp"],["jdjdb","azdwg"],["xvn","ztp"],[]]`

	arr := strings.Split(str, "],")


	fmt.Println(arr[31])
	//fmt.Println(arr[47])
	//fmt.Println(arr[48])

	fmt.Println(arr[33])
	fmt.Println(arr[34])
}


/*



["ohmvnkxnokpnwxmr","jdjdb","azdwg","xvn","mljb","rdaavykiextpr","ztp"]
[["azdwg"],["ohmvnkxnokpnwxmr","mljb"],["jdjdb"],[],[],[],["xvn","mljb"],[],["xvn"],[],["xvn"],["jdjdb","mljb"],[],["ohmvnkxnokpnwxmr","mljb","ztp"],["ohmvnkxnokpnwxmr"],[],[],["ohmvnkxnokpnwxmr","jdjdb"],["jdjdb"],[],[],[],["azdwg","mljb"],["ztp"],["xvn"],["jdjdb"],["mljb"],["xvn"],["rdaavykiextpr"],[],[],["ohmvnkxnokpnwxmr","jdjdb","mljb","rdaavykiextpr"],[],["jdjdb","mljb","rdaavykiextpr","ztp"],["ohmvnkxnokpnwxmr"],["rdaavykiextpr","ztp"],["mljb"],["rdaavykiextpr"],[],["azdwg"],[],["rdaavykiextpr"],["jdjdb"],[],["mljb","ztp"],[],["rdaavykiextpr","ztp"],["jdjdb","azdwg"],["xvn","ztp"],[]]
 */


// 使用channel实现互斥锁
func Channel() {
	var count int
	flag := make(chan struct{}, 1)
	group := sync.WaitGroup{}
	for i:=0; i<10000; i++{
		group.Add(1)
		go func() {
			flag <- struct{}{}
			count ++
			<-flag
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(count)
}

// 使用mutex
func Mutex() {
	var count int

	flag := sync.Mutex{}

	group := sync.WaitGroup{}
	for i:=0; i<10000; i++{
		group.Add(1)
		go func() {
			flag.Lock()
			count ++
			flag.Unlock()
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(count)

}