/**
 *@Author luojunying
 *@Date 2022-06-29 22:02
 */
package main

import (
	"errors"
	"fmt"
)

type ListStack struct {
	list    []int
	minList []int
	index   int
}

func NewListStack(size int) *ListStack {
	return &ListStack{
		list:    make([]int, size, size),
		minList: make([]int, size, size),
		index:   -1,
	}
}

func (l *ListStack) Push(key int) error {
	if l.index == len(l.list)-1 {
		return errors.New("list is full")
	}
	k, err := l.getMin()
	if err != nil {
		l.minList[l.index+1] = key
	} else if key >= k {
		l.minList[l.index+1] = k
	} else {
		l.minList[l.index+1] = key
	}
	l.index++
	l.list[l.index] = key
	return nil
}

func (l *ListStack) Pop() (int, error) {
	if l.index == -1 {
		return 0, errors.New("list is empty")
	}
	k := l.list[l.index]
	l.index--
	return k, nil
}

func (l *ListStack) getMin() (int, error) {
	if l.index == -1 {
		return 0, errors.New("list is empty")
	}
	return l.minList[l.index], nil
}

func main() {
	l := NewListStack(5)
	fmt.Println(*l)
	err := l.Push(1)
	fmt.Println("err: ", err)
	err = l.Push(2)
	fmt.Println("err: ", err)
	err = l.Push(3)
	fmt.Println("err: ", err)
	err = l.Push(4)
	fmt.Println("err: ", err)
	m, err := l.getMin()
	fmt.Println("min: ", m, " err:", err)

	err = l.Push(5)
	fmt.Println("err: ", err)
	err = l.Push(6)
	fmt.Println("err: ", err)

	k, err := l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)
	err = l.Push(-6)
	fmt.Println("err: ", err)

	fmt.Println(l.getMin())
	err = l.Push(6)
	fmt.Println("err: ", err)

	k, err = l.Pop()
	fmt.Println("k: ", k, " err: ", err)

	fmt.Println(*l)
}
