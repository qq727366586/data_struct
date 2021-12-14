/**
 *@Author luojunying
 *@Date 2021-12-14 21:08
 */
package main

import "fmt"

type Node struct {
	Value int
	next *Node
}

func (n *Node) Add(temp Node) *Node {
	var node = n
	for node.next != nil {
		node = node.next
	}
	node.next = &temp
	return n
}

func (n *Node) List() {
	for n != nil {
		fmt.Println(n.Value)
		n = n.next
	}
}

func (n *Node) reverse() *Node {
	var pre *Node
	var node = n
	for node != nil {
		temp := node.next
		node.next = pre
		pre = node
		node = temp
	}
	return pre
}

func main() {
	node := &Node{
		Value: 1,
		next: nil,
	}
	node.Add(Node{Value: 2})
	node.Add(Node{Value: 3})
	node.Add(Node{Value: 4})
	node.Add(Node{Value: 5})
	node.Add(Node{Value: 6})
	node.List()
	node = node.reverse()

	node.List()
}