package main

import "fmt"

//定义结构体
type Node struct {
	Data int
	Next *Node
}

func NewNode() *Node {
	return &Node{}
}

// 无序插入
func (head *Node) Insert(node *Node) *Node {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = node
	return head
}

// 有序插入
func (head *Node) SortInsert(node *Node) *Node {
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Data >= node.Data {
			break
		}
		tmp = tmp.Next
	}
	next := tmp.Next
	tmp.Next = node
	node.Next = next
	return head
}

func main() {
	Test()
}


func Test() {
	head := NewNode()
	head = head.Insert(&Node{Data: 21}).Insert(&Node{Data: 100, Next: nil}).Insert(&Node{Data: 100}).Insert(&Node{Data: 100}).Insert(&Node{Data: 100}).Insert(&Node{Data: 100})
	for head.Next != nil {
		fmt.Println(head.Next.Data)
		head = head.Next
	}
	//
	head = head.SortInsert(&Node{Data: 21}).SortInsert(&Node{Data: 131}).SortInsert(&Node{Data: 10}).SortInsert(&Node{Data: -1}).SortInsert(&Node{Data: -91})
	for head.Next != nil {
		fmt.Println(head.Next.Data)
		head = head.Next
	}
}

