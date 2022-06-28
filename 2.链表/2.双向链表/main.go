/**
 *@Author luojunying
 *@Date 2022-01-15 16:28
 */
package main

import "fmt"

type Node struct {
	data int
	pre  *Node
	next *Node
}
type List struct {
	head *Node
}

func NewList() *List {
	return &List{}
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
	node.pre = head
	return l
}

// 有序插入
func (l *List) InsertSort(node *Node) *List {
	if l.head == nil {
		l.head = node
		return l
	}
	head := l.head
	var pre *Node
	for head != nil {
		pre = head
		if head.data >= node.data {
			head.data, node.data = node.data, head.data
			next := head.next
			head.next = node
			node.pre = head
			node.next = next
			if next != nil {
				next.pre = node
			}
			break
		}
		head = head.next
	}
	pre.next = node
	node.pre = pre
	return l
}

func (l *List) Remove(key int) *List {
	if l.head == nil {
		return l
	}
	// 看是否第一个元素
	if l.head.data == key {
		tmp := l.head.next
		l.head.next = nil
		l.head = tmp
	}
	head := l.head
	for head != nil && head.next != nil {
		if head.next.data == key {
			tmp := head.next.next
			head.next.pre = nil
			head.next.next = nil
			head.next = tmp
			if tmp != nil {
				head.next.pre = head
			}
			break
		}
		head = head.next
	}
	return l
}

// 正向遍历
func (l *List) ForwardPrint() {
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

// 反向遍历
func (l *List) ReversePrint() {
	node := l.head
	for node.next != nil {
		node = node.next
	}
	for node != nil {
		if node.pre != nil {
			fmt.Print(node.data, ",")
		} else {
			fmt.Print(node.data)
		}
		node = node.pre
	}
}

func main() {
	list := NewList()
	l := list.InsertSort(&Node{data: 1}).InsertSort(&Node{data: 11}).InsertSort(&Node{data: 1}).InsertSort(&Node{data: 111}).InsertSort(&Node{data: -1})
	//l.Remove(1).Remove(111).Remove(-1).Remove(1).Remove(11).Remove(123123)
	l.ForwardPrint()
}
