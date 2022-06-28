/**
 *@Author luojunying
 *@Date 2022-06-23 20:53
 */
package main

import "fmt"

//定义结构体
type Node struct {
	data int
	next *Node
	pre  *Node
}

type LinkedStack struct {
	head *Node
	end  *Node
}

func (l *LinkedStack) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedStack) Size() int {
	count := 0
	tmp := l.head
	for tmp != nil {
		count++
		tmp = tmp.next
	}
	return count
}

func (l *LinkedStack) Push(n *Node) *LinkedStack {
	if l.head == nil {
		l.head, l.end = n, n
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

func (l *LinkedStack) Pop() *Node {
	if l.end == nil {
		return l.end
	}
	tmp := l.end
	l.end = l.end.pre
	if l.end != nil {
		l.end.next.pre = nil
		l.end.next = nil
	} else {
		l.head = l.end
	}
	return tmp
}

func (l *LinkedStack) Top() *Node {
	return l.end
}

func main() {
	stack := &LinkedStack{}
	stack.Push(&Node{data: 1}).Push(&Node{data: 11}).Push(&Node{data: 1}).Push(&Node{data: 111})
	fmt.Println("size ", stack.Size(), "isEmpty ", stack.IsEmpty())
	fmt.Println("pop ", stack.Pop(), "pop ", stack.Pop(), "pop ", stack.Pop(), "pop ", stack.Pop(), "pop ", stack.Pop())
	fmt.Println("size ", stack.Size(), "isEmpty ", stack.IsEmpty())
}
