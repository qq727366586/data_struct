/**
 *@Author luojunying
 *@Date 2022-01-15 16:28
 */
package main

import "fmt"

type man struct {
	no   int
	name string
	pre  *man
	next *man
}

//插入
func (m *man) insert(newMan *man) *man {
	temp := m
	if temp.next != nil {
		temp = temp.next
	}
	temp.next = newMan
	newMan.pre = temp
	return m
}

func (m *man) reverse(headMan *man) *man {
	var pre *man
	var next *man
	for headMan != nil {
		next = headMan.next
		headMan.next = pre
		headMan.pre = next
		pre = headMan
		headMan = next
	}
	return headMan
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
	newMan.pre = temp
	if newMan.next != nil {
		newMan.next.pre = newMan
	}
	return m
}

//删除
func (m *man) delete(no int) *man {
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

		if temp.next != nil {
			temp.next.pre = temp
		}
	}
	return m
}

func (m *man) list() {
	//先正后反遍历
	temp := m
	if temp.next == nil {
		return
	}
	for temp.next != nil {
		fmt.Printf(temp.name + ", ")
		temp = temp.next
	}
	fmt.Printf(temp.name + ", ")
	fmt.Println()
	for temp.pre != nil {
		fmt.Printf(temp.name + ", ")
		temp = temp.pre
	}
	fmt.Printf(temp.name + ", ")
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
	root6 := &man{
		no: 6,
		name: "机器人6号",
	}
	root.insertSort(root1).insertSort(root6).insertSort(root2).insertSort(root5).insertSort(root4).insertSort(root3).insertSort(root4)
	root.delete(1)
	root.list()
}
