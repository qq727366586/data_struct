package main

import "fmt"

//定义结构体
type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func NewList() *List {
	return &List{
		head: nil,
	}
}

// 无序插入
func (l *List) Insert(node *Node) *List {
	if l.head == nil {
		l.head = node
		return l
	}
	head := l.head
	for head.next != nil {
		head = head.next
	}
	head.next = node
	return l
}

// 有序插入
func (l *List) InsertSort(node *Node) *List {
	if l.head == nil {
		l.head = node
		return l
	}
	head := l.head
	var pre, next *Node
	for head != nil {
		pre = head
		if head.data >= node.data {
			tmp := node.data
			node.data = head.data
			head.data = tmp
			next = head.next
			head.next = node
			node.next = next
			break
		}
		head = head.next
	}
	pre.next = node
	return l
}

func (l *List) Remove(key int) *List {
	head := l.head
	if head == nil {
		return l
	}
	// 判断第一个元素
	if head.data == key {
		tmp := head.next
		head.next = nil
		l.head = tmp
	}
	for head != nil && head.next != nil {
		if head.next.data == key {
			tmp := head.next
			head.next = nil
			head.next = tmp.next
			break
		}
		head = head.next
	}
	return l
}

// 反转
func (l *List) Reverse() *List {
	head := l.head
	var pre, next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	l.head = pre
	return l
}

// 遍历
func (l *List) Print() {
	node := l.head
	for node != nil {
		if node.next != nil {
			fmt.Print(node.data, ",")
		} else {
			fmt.Print(node.data)
		}
		node = node.next
	}
}

func main() {
	list := NewList()
	list.Reverse()
	l := list.InsertSort(&Node{data: 1}).InsertSort(&Node{data: 11}).InsertSort(&Node{data: 1}).InsertSort(&Node{data: 111}).InsertSort(&Node{data: -1})
	l.Reverse()
	l.Remove(11).Remove(1).Remove(-1).Remove(111)
	l.Print()
}
