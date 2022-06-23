/**
 *@Author luojunying
 *@Date 2022-06-23 20:53
 */
package main

import (
	"errors"
	"fmt"
)

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

func (l *LinkedStack) Push(data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head, l.end = node, node
		return
	}
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
	node.pre = tmp
	l.end = tmp.next
}

func (l *LinkedStack) Pop() (int, error) {
	if l.end == nil {
		return 0, errors.New("stack no data")
	}
	out := l.end.data
	// 说明只有一个元素
	if l.end.pre == nil {
		l.head = nil
		l.end = nil
		return out, nil
	}
	l.end = l.end.pre
	l.end.next.pre = nil
	l.end.next = nil
	return out, nil
}

func (l *LinkedStack) Top() (int, error) {
	if l.end != nil {
		return l.end.data, nil
	}
	return 0, errors.New("stack no data")
}

func main() {
	stack := &LinkedStack{}
	stack.Push(1)
	data, err := stack.Top()
	fmt.Println(data, err, "top") // 1 nil

	stack.Push(2)
	data, err = stack.Top()
	fmt.Println(data, err, "top") // 2 nil

	num, err := stack.Pop()
	fmt.Println(num, err) // 2 nil

	data, err = stack.Top()
	fmt.Println(data, err, "top") // 1 nil

	num, err = stack.Pop()
	fmt.Println(num, err) // 1 nil

	data, err = stack.Top()
	fmt.Println(data, err, "top") // 0 err

	num, err = stack.Pop() // 0 err
	fmt.Println(num, err)

	stack.Push(2)
	fmt.Println(stack.Size())    // 1
	fmt.Println(stack.IsEmpty()) // false
	num, err = stack.Pop()
	fmt.Println(num, err)
	fmt.Println(stack.Size())
	fmt.Println(stack.IsEmpty())
}
