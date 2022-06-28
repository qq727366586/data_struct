/**
 *@Author luojunying
 *@Date 2022-06-28 22:04
 */
package main

import "fmt"

// 实现双端队列

//定义结构体
type Node struct {
	data int
	next *Node
	pre  *Node
}

type LinkedQueue struct {
	head *Node
	end  *Node
}

func (l *LinkedQueue) pushHead(n *Node) *LinkedQueue {
	if l.head == nil {
		l.head = n
		l.end = n
		return l
	}
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = n
	n.pre = tmp
	l.end = n
	return l
}

func (l *LinkedQueue) pushTail(n *Node) *LinkedQueue {
	if l.end == nil {
		l.head = n
		l.end = n
		return l
	}
	l.end.next = n
	n.pre = l.end
	l.end = n
	return l
}

func (l *LinkedQueue) Pop() *Node {
	if l.head == nil {
		return l.head
	}
	tmp := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.pre.next = nil
		l.head.pre = nil
	} else {
		l.end = nil
	}
	return tmp
}

func (l *LinkedQueue) Size() int {
	var count = 0
	tmp := l.head
	for tmp != nil {
		count++
		tmp = tmp.next
	}
	return count
}

func (l *LinkedQueue) IsEmpty() bool {
	return l.head == l.end
}

func main() {
	l := &LinkedQueue{}
	l.pushHead(&Node{data: 1}).pushHead(&Node{data: 2})
	fmt.Println("size ", l.Size())
	fmt.Println("isEmpty ", l.IsEmpty())

	l.pushTail(&Node{data: 10})
	fmt.Println("size ", l.Size())

	fmt.Println("pop ", l.Pop())
	fmt.Println("pop", l.Pop())
	fmt.Println("pop", l.Pop())
	fmt.Println("pop", l.Pop())
	fmt.Println("pop", l.Pop())
	fmt.Println("size ", l.Size())
	fmt.Println("isEmpty ", l.IsEmpty())
	fmt.Println(l.head, l.end)
}
