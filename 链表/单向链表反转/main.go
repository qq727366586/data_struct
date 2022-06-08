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

func (n *Node) Add(node *Node) {
	n.next = node
}

func (n *Node) Reverse(head *Node) *Node {
	var pre *Node
	for head != nil {
		next := head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	n := &Node{
		1, nil,
	}
	n.Add(&Node{2, nil})
	head := n.Reverse(n)
	fmt.Println(head, head.next)
}