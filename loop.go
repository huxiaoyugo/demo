package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)




type  FuncDo func()


func(f FuncDo) Do() {
	f()
}

type I interface {
	Do()
}


type S struct  {
	*T
}

type T struct {

}


func (t* T) Do() {

}

func test() {
	fmt.Println("test")
}
func main() {

	ch  := make(chan int, 0)
	go func() {
		time.Sleep(3*time.Second)
		ch <- 1
	}()

	select {
	case v := <- ch:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}


	fmt.Println("end")



}









var chMap = sync.Map{} // 全局的map,保存请求的chan
func Loop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()// 获取参数
	uid, _     := strconv.Atoi(r.FormValue("uid"))
	//lastId, _     := strconv.Atoi(r.FormValue("lastId"))
	// 删除消费过的数据,删除<lastId的数据, 检查redis中是否有历史数据,如果有的话,直接返回
	// 如果没有的话
	ch := make(chan []byte)
	chMap.Store(uid,ch)
	select {
		case val := <- ch:
			fmt.Println(val)
			// 写返回的数据等操作
			w.Write(val)
		case <-time.After(30*time.Second):
			fmt.Println("超时")
	}
	// 删除ch
	chMap.Delete(uid)
	return
}
func sub() {
	// redis.sub("key")  订阅相关消息, 比如支付完成，从订阅中获得的数据，读取uid
	msg := "{uid:10000006,orderId:1009204029,msgId:1023,type:101}"
	uid := 10000006
	if ch, ok := chMap.Load(uid); ok {
		// 发送消息
		ch.(chan []byte) <- []byte(msg)
	}
	// 如果没有连接，不做处理
	return
}
