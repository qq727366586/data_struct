/**
 *@Author luojunying
 *@Date 2022-01-15 15:15
 */
package main

import (
	"fmt"
)

type man struct {
	no   int
	name string
	next *man
}

//普通插入
func (m *man) insert(newMan *man) *man {
	temp := m
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newMan
	return m
}

//有序插入
func (m *man) insertSort(newMan *man) *man {
	temp := m
	for temp.next != nil {
		if temp.next.no > newMan.no {
			break
		}
		if temp.next.no == newMan.no {
			return m
		}
		temp = temp.next
	}
	newMan.next = temp.next
	temp.next = newMan
	return m
}

//删除
func (m *man) delete(no int) {
	temp := m
	boolean := false
	for temp.next != nil {
		if temp.next.no == no {
			boolean = true
			break
		}
		temp = temp.next
	}
	if boolean {
		temp.next = temp.next.next
	}
}

//遍历
func (m *man) list() {
	temp := m
	//先判断是否为空链表
	if temp.next == nil {
		return
	}
	for temp.next != nil {
		fmt.Printf(temp.next.name + ",  ")
		temp = temp.next
	}
}

func main() {
	root := &man{
		no: 0,
		name: "机器人头号",
	}

	root1 := &man{
		no: 1,
		name: "机器人1号",
	}

	root2 := &man{
		no: 2,
		name: "机器人2号",
	}
	root4 := &man{
		no: 4,
		name: "机器人4号",
	}
	root3 := &man{
		no: 3,
		name: "机器人3号",
	}
	root5 := &man{
		no: 5,
		name: "机器人5号",
	}
	root.insertSort(root1).insertSort(root2).insertSort(root5).insertSort(root4).insertSort(root3).insertSort(root4)
	root.delete(6)
	root.list()
}