package 工作中的问题

import "fmt"

type Node struct {
	Val *maintainInfo
	Next *Node
}

type maintainInfo struct {
	MachineId int
	SiteId int
	CityId int
	BeginTime int64
	EndTime int64
}


type Head struct {
	// 节点数
	NodeCount int
	TotalVal int
	Next *Node
}



func CreateLinkMap(maintainList []*maintainInfo) map[int]*Head {

	var LinkMap = make(map[int]*Head)

	for _, maintainInfo := range maintainList {
		head := LinkMap[maintainInfo.MachineId]
		if head == nil {
			head = &Head{}
			LinkMap[maintainInfo.MachineId] = head
		}

		newNode := &Node{
			Val: maintainInfo,
			Next:nil,
		}
		// 向head中插入一个新的节点
		InsertIntoHead(head, newNode)
	}
	return LinkMap
}


func InsertIntoHead(head *Head, node *Node) {

	if head.NodeCount == 0 {
		head.Next = node
		head.NodeCount +=1
		return
	}

	var last *Node
	// 1、找到合适的位置插入
	p := head.Next
	for {
		if p == nil {
			last.Next = node
			break
		}
		if node.Val.BeginTime > p.Val.BeginTime {
			last = p
			p = p.Next
		} else {
			node.Next = p
			if last == nil {
				head.Next = node
			} else {
				last.Next = node
			}
			break
		}
	}
	mergeNode(head)
}


func mergeNode(head *Head) {
	if head.NodeCount == 0 {
		return
	}
	p := head.Next
	for {
		next := p.Next
		if next == nil {
			return
		}

		if next.Val.BeginTime > p.Val.EndTime {
			p = next
			continue
		}

		// 合并节点p和next

		maxEnd := next.Val.EndTime
		if p.Val.EndTime > maxEnd {
			maxEnd = p.Val.EndTime
		}

		p.Val.EndTime = maxEnd
		// 删除next节点
		p.Next = next.Next
	}
}


func GetRangeVal(head *Head)int {
	rangeVal := 0
	for p := head.Next; p != nil; p=p.Next {
		rangeVal += int(p.Val.EndTime-p.Val.BeginTime+1)
	}
	return rangeVal
}

func PrintHead(head *Head) {

	p := head.Next
	for ;p != nil; p= p.Next {
		fmt.Printf("[%d,%d]->", p.Val.BeginTime, p.Val.EndTime)
	}
	fmt.Println("")
}


func Merge(maintainList []*maintainInfo) (res []*maintainInfo) {
	headMap := CreateLinkMap(maintainList)
	for _, head := range headMap {
		for p:=head.Next; p != nil; p=p.Next {
			res = append(res, p.Val)
		}
	}
	return
}