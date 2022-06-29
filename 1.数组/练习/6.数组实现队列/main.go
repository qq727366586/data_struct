/**
 *@Author luojunying
 *@Date 2022-06-29 22:35
 */
package main

import (
	"errors"
	"fmt"
)

type listQueue struct {
	list  []int
	size  int
	begin int
	end   int
}

func NewListQueue(size int) *listQueue {
	return &listQueue{
		list:  make([]int, size, size),
		size:  0,
		begin: 0,
		end:   0,
	}
}

func (l *listQueue) push(k int) error {
	if l.size == len(l.list) {
		return errors.New("list is full")
	}
	l.size++
	l.list[l.end] = k
	if l.end < len(l.list)-1 {
		l.end++
	} else {
		l.end = 0
	}
	return nil
}

func (l *listQueue) pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("list is empty")
	}
	l.size--
	ans := l.list[l.begin]
	if l.begin < len(l.list)-1 {
		l.begin++
	} else {
		l.begin = 0
	}
	return ans, nil
}

func main() {
	l := NewListQueue(5)
	err := l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)
	err = l.push(1)
	fmt.Println("err: ", err)

	k, err := l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)
	k, err = l.pop()
	fmt.Println("k: ", k, " err: ", err)

}
