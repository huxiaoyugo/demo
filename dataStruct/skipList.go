package main

import (
	"math/rand"
	"time"
	"fmt"
)


/*
跳跃表
2019-05-17
*/
const (
	MAX_LEVEL = 5
)
func main() {
	list := NewSkipList()
	for i  :=0; i<50000; i++ {
		list.Insert(fmt.Sprintf("member%d",i),int64(i))
	}


	tmStart :=time.Now()

	no := list.Find(0)
	fmt.Println("Find:",time.Now().Sub(tmStart))

	if no == nil {
		fmt.Println("result is nil")
	} else {
		fmt.Printf("score:%d, member:%v\n",no.Score,no.Member)
	}



	tmStart =time.Now()

	list.Delete(0)

	fmt.Println("Delte:",time.Now().Sub(tmStart))

	tmStart =time.Now()
	no = list.Find(0)

	fmt.Println("Find:",time.Now().Sub(tmStart))
	if no == nil {
		fmt.Println("result is nil")
	} else {
		fmt.Printf("score:%d, member:%v\n",no.Score,no.Member)
	}


}


type skipListLevel struct {
	NextNode *skipListNode
}

type skipListNode struct {
	Member interface{}
	Score  int64
	Level []*skipListLevel
	PrevNode *skipListNode
}


type skipList struct {
	Header *skipListNode
	Tail *skipListNode
	MaxLevel  int
	Len  int
}


func NewSkipList() *skipList {
	skipList := skipList{
		MaxLevel:0,
		Len:0,
		Tail:nil,
	}

	// 头部节点不做为存储数据
	// level直接设为为32
	skipList.Header = createSkipListNode(nil,0,MAX_LEVEL)
	return &skipList
}

func(list *skipList) Insert(data interface{}, score int64) {

	node := createSkipListNode(data, score,0)

	preNode := list.findInsetPreNode(score)


	if len(node.Level) > list.MaxLevel {
		list.MaxLevel = len(node.Level) - 1
	}
	list.Len++

	// 先插入第0层
	node.PrevNode = preNode
	node.Level[0].NextNode = preNode.Level[0].NextNode

	if preNode.Level[0].NextNode != nil {
		preNode.Level[0].NextNode.PrevNode = node
	} else { // 为空说明是最后一个节点了，所以
		list.Tail = node
	}
	preNode.Level[0].NextNode = node

	// 处理索引层
	for levelIndex := 1; levelIndex < len(node.Level); {
		backNode := preNode
		// preNode开始向后搜索，知道找到一个拥有这个层的节点，并记录向回搜索的长度
		for {
			if len(backNode.Level) > levelIndex {
				// 符合要求
				break
			} else { // 没有该层，继续往回搜索
				backNode = backNode.PrevNode
			}
		}

		// 处理
		for ; len(backNode.Level) > levelIndex && levelIndex < len(node.Level); levelIndex++ {
			node.Level[levelIndex].NextNode = backNode.Level[levelIndex].NextNode
			backNode.Level[levelIndex].NextNode = node
		}
		backNode = backNode.PrevNode
	}
}


func(list *skipList) Find(score int64) *skipListNode {

	curNode := list.Header
	curLevel := list.MaxLevel

	for {
		if curNode.Level[curLevel].NextNode == nil ||curNode.Level[curLevel].NextNode.Score > score{
			if curLevel == 0 {
				return nil
			} else {
				// 检查下一层
				curLevel-=1
				continue
			}
		}
		if curNode.Level[curLevel].NextNode.Score < score {
			curNode = curNode.Level[curLevel].NextNode
			continue
		} else {// 相等 只需要找到这个节点的前一个节点即可
			return curNode.Level[curLevel].NextNode
		}
	}
}


func(list *skipList) Delete(score int64) {

	node := list.Find(score)
	if node == nil {
		return
	}

	// 先处理第0层
	node.PrevNode.Level[0].NextNode = node.Level[0].NextNode
	if node.Level[0].NextNode == nil {
		// 更换尾部节点
		list.Tail = node.PrevNode
	}

	backNode := node.PrevNode
	for curLevel := 1; curLevel <len(node.Level); {

		for {
			if len(backNode.Level) > curLevel {
				break
			} else {
				backNode = backNode.PrevNode
			}
		}

		for ; curLevel < len(backNode.Level) && curLevel <len(node.Level); curLevel++ {
			backNode.Level[curLevel].NextNode = node.Level[curLevel].NextNode
		}
		backNode = backNode.PrevNode
	}
}

// 查找出待插入的
func(list *skipList) findInsetPreNode(score int64) *skipListNode {

	if list.Len == 0 {
		return list.Header
	}

	curNode := list.Header
	curLevel := list.MaxLevel

	for {
		 if curNode.Level[curLevel].NextNode == nil ||curNode.Level[curLevel].NextNode.Score > score{
			if curLevel == 0 {
				return curNode
			} else {
				// 检查下一层
				curLevel-=1
				continue
			}
		}
		if curNode.Level[curLevel].NextNode.Score < score {
			curNode = curNode.Level[curLevel].NextNode
			continue
		} else {// 相等 只需要找到这个节点的前一个节点即可
			return curNode.Level[curLevel].NextNode.PrevNode
		}
	}
}



func createSkipListNode(data interface{}, score int64, fixedLevel int ) (*skipListNode) {
	level := 0
	if fixedLevel <=0 {
		rand.Seed(time.Now().UnixNano())
		level = rand.Int() % MAX_LEVEL
		fixedLevel = level+1
	} else {
		level = fixedLevel-1
	}
	node := skipListNode {
		Member: data,
		Score: score,
		PrevNode:nil,
	}
	node.Level = make([]*skipListLevel, fixedLevel)
	for i := 0; i <= level; i++ {
		node.Level[i] = &skipListLevel{}
	}
	return &node
}


func (list *skipList) ShowLevel(level int) {
	if list.MaxLevel < level {
		return
	}
	node := list.Header.Level[level].NextNode
	for {
		if node == nil {
			break
		}
		fmt.Printf("%5d", node.Score)
		node = node.Level[level].NextNode
	}
	fmt.Printf("\n")
}

func (list *skipList)ShowAllLevel() {
	for level := MAX_LEVEL-1 ;level>=0; level-- {
		list.ShowLevel(level)
	}
}