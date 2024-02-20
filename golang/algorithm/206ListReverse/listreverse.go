package main

import (
	"fmt"

	"example.com/m/v2/algorithm/list"
)

func reverse(l *list.ListNode) *list.ListNode {
	head := new(list.ListNode)
	head.Next = nil
	node := new(list.ListNode)
	for l != nil {
		node = l
		l = l.Next
		node.Next = head.Next
		head.Next = node
	}
	return head.Next
}

func main() {
	l := list.CreateNewList()
	bak := l
	for bak != nil {
		fmt.Println(bak.Val)
		bak = bak.Next
	}
	fmt.Println("--------------")
	l = reverse(l)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
