/**
 *@Author luojunying
 *@Date 2022-01-05 20:12
 */
package main

import "fmt"

//前缀树

type prefixTree struct {
	pass int
	end  int
	next []*prefixTree
}

func NewPrefixTree() *prefixTree {
	var next [26]*prefixTree
	return &prefixTree{
		pass: 0,
		end:  0,
		next: next[:],
	}
}

func Insert(prefixTree *prefixTree, str string) {
	root := prefixTree
	root.pass++
	//将字符串转字符类型数组
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		//如果为空,则添加
		if root.next[index] == nil {
			root.next[index] = NewPrefixTree()
		}
		root = root.next[index]
		root.pass++
	}
	root.end++
}

//查询单词出现过几次
func search(prefixTree *prefixTree, str string) int {
	if str == "" {
		return 0
	}
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		if prefixTree.next[index] == nil {
			return 0
		}
		prefixTree = prefixTree.next[index]
	}
	return prefixTree.end
}

//出现前缀几次
func findPreNumber(prefixTree *prefixTree, str string) int {
	if str == "" {
		return 0
	}
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		if prefixTree.next[index] == nil {
			return 0
		}
		prefixTree = prefixTree.next[index]
	}
	return prefixTree.pass
}

//删除某个节点
func delete(prefixTree *prefixTree, str string) {
	if search(prefixTree ,str) != 0 {
		bt := []byte(str)
		root := prefixTree
		for _, v := range bt {
			index := v - 'a'
			if root.next[index].pass == 0 {
				root.next[index] = nil
				return
			}
			root = root.next[index]
		}
		root.end--
	}
}

func main() {
	str := []string{"aa", "ab", "abc", "cd", "cd"}
	p := NewPrefixTree()
	for _, v:= range str {
		Insert(p ,v)
	}
	end := findPreNumber(p, "ab")
	delete(p, "ab")
	end = search(p, "ab")
	fmt.Println(end)
}
