/**
 *@Author luojunying
 *@Date 2021-11-28 17:55
 */
package main

import (
	"fmt"
)

//思路,先把顶级找出,在找子

type Menu struct {
	Id       int
	Pid      int
	Name     string
	Children []Menu
}


func GetMenu(menu []Menu, pid int) []Menu {
	var out []Menu
	for _, v := range menu {
		if v.Pid == pid {
			child := GetMenu(menu, v.Id)
			node := Menu{
				Id:   v.Id,
				Pid:  v.Pid,
				Name: v.Name,
			}
			node.Children = child
			out = append(out, node)
		}
	}
	return out
}

//查询自己 和其子节点 所有id  一般用于删除
func FindTreeChildNode(arr []Menu, id int, temp []int) []int {
	for _, v := range arr {
		if v.Pid == id {
			temp = append(temp, v.Id)
			temp = FindTreeChildNode(arr, v.Id, temp)
		} else if v.Id == id && v.Pid == 0 {
			temp = append(temp, v.Id)
		}
	}
	return temp
}

func main() {
	menu := []Menu{
		{Id: 1, Pid: 2, Name: "可爱"},
		{Id: 2, Pid: 0, Name: "形容词"}, // 1
		{Id: 3, Pid: 2, Name: "丑"},
		{Id: 4, Pid: 0, Name: "机器人"}, // 1
		{Id: 5, Pid: 4, Name: "太空机器人"},
		{Id: 6, Pid: 4, Name: "地球机器人"},
		{Id: 7, Pid: 5, Name: "太空机器人-小傻"},
		{Id: 8, Pid: 5, Name: "太空机器人-小思"},
		{Id: 9, Pid: 8, Name: "太空机器人-小思-儿子"},
	}
	out := GetMenu(menu, 0)
	fmt.Print(out)
	fmt.Println()
	var temp []int
	nodes := FindTreeChildNode(menu, 2, temp)
	fmt.Println(nodes)
}
