package main

import "fmt"

func main() {
	//l := []int{1,2,3,4,5,6,7}

	al := Constructor()
	al.Inc("a")
	//al.Inc("b")
	//al.Inc("b")
	//al.Inc("b")
	al.Inc("b")
	al.Inc("c")
	al.Inc("d")
	al.Inc("e")
	al.Inc("c")
	//al.Inc("a")
	al.desc()

}
func (this *AllOne)desc () {
	p := this.head
	if p != nil {
		fmt.Printf("head: %s: %d\n", p.Key, p.Val)
	}
	p = this.tail
	if p != nil {
		fmt.Printf("tail: %s: %d\n", p.Key, p.Val)
	}
	p = this.head
	for p != nil {
		fmt.Printf("%s: %d, ", p.Key, p.Val)
		p = p.Next
	}
	fmt.Println()
}

type Node struct {
	Key string
	Val int
	Next *Node
	Pre *Node
}


type AllOne struct {
	head *Node
	tail *Node
	m map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		m : make(map[string]*Node),
	}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	node := this.m[key]
	if node != nil {
		node.Val += 1
		if node.Next != nil && node.Next.Val < node.Val {
			t := node.Next

			if node.Pre != nil {
				node.Pre.Next = t
			}
			t.Pre = node.Pre
			node.Next = t.Next
			if node.Next != nil {
				node.Next.Pre = node
			}
			node.Pre = t
			t.Next = node

			if node.Next == nil {
				this.tail = node
			}

			if t.Pre == nil {
				this.head = t
			}
		}
	} else {
		node := &Node{
			Key : key,
			Val : 1,
			Next: this.head,
		}
		this.head = node
		if this.head.Next != nil {
			this.head.Next.Pre = this.head
		} else {
			this.tail = node
		}
		this.m[key] = node
	}
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	node := this.m[key]
	if node == nil {
		return
	}

	if node.Val == 1 {
		// 删除
		delete(this.m, key)
		// 删除的是第一个节点
		if node.Pre == nil {
			this.head = node.Next
			if this.head == nil {
				this.tail = nil
			} else {
				this.head.Pre = nil
			}
			return
		}
		// 删除的是最后一个节点
		if node.Next == nil {
			node.Pre.Next = nil
			this.tail = node.Pre
			return
		}
		// 删除中间节点
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		return
	}


	// 减一
	node.Val -= 1

	if node.Pre != nil && node.Pre.Val > node.Val {
		// 交换
		t := node
		node = node.Pre
		if node.Pre != nil {
			node.Pre.Next = t
		}
		t.Pre = node.Pre
		node.Next = t.Next
		if node.Next.Pre != nil {
			node.Next.Pre = node
		}
		node.Pre = t
		t.Next = node

		if t.Pre == nil {
			this.head = t
		}

		if node.Next == nil {
			this.tail = node
		}
	}
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail != nil {
		return this.tail.Key
	}
	return ""
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head != nil {
		return this.head.Key
	}
	return ""
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */