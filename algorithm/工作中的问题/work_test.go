package 工作中的问题

import (
	"testing"
	"fmt"
	"sync/atomic"
)

func TestCreateLinkMap(t *testing.T) {

	var maintainList = make([]*maintainInfo,0)
	maintainList = append(maintainList,

		&maintainInfo{
			MachineId:1,
			BeginTime: 5,
			EndTime:6,
		},

		&maintainInfo{
			MachineId:1,
			BeginTime: 7,
			EndTime:7,
		},

		&maintainInfo{
			MachineId:2,
			BeginTime:2,
			EndTime:10,
		},

		&maintainInfo{
			MachineId:2,
			BeginTime: 9,
			EndTime:15,
		},
		&maintainInfo{
			MachineId:2,
			BeginTime: 9,
			EndTime:15,
		},

		&maintainInfo{
			MachineId:2,
			BeginTime: 88,
			EndTime:95,
		},
		&maintainInfo{
			MachineId:2,
			BeginTime: 8,
			EndTime:130,
		},
	)

	maintainList = Merge(maintainList)
	//PrintHead(head)

	for _, item := range maintainList {

		fmt.Printf("[%d,%d],", item.BeginTime, item.EndTime)

	}


}

var (
	finish = make(chan string, 1)

	count int32 = 0
)

func Test_Channel(t *testing.T) {


	go func() {
		for {
			select {
			case val := <-finish:
				fmt.Println("1:",val)
			}
		}
	}()

	go func() {
		for {
			select {
			case val := <-finish:
				fmt.Println(val)
			}
		}
	}()

	atomic.AddInt32(&count, 1)
	defer func() {
		atomic.AddInt32(&count, -1)
		finish<-"ddd"
		finish<-"fee"
		finish<-"gg"
		finish<-"fehhe"
	}()

}