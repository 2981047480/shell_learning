package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateNewList() *ListNode {
	head := new(ListNode)
	now := head
	for i := 0; i < 10; i++ {
		node := new(ListNode)
		node.Val = i
		now.Next = node
		now = now.Next
	}
	now.Next = nil
	return head.Next
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var smalltail = new(ListNode)
	var bigtail = new(ListNode)
	smallhead := smalltail // 新建链表的头
	bighead := bigtail     // 现有链表的头
	for now := head; now != nil; now = now.Next {
		if now.Val < x {
			smalltail.Next = now
			smalltail = smalltail.Next
		} else {
			bigtail.Next = now
			bigtail = bigtail.Next
		}
	}
	smalltail.Next = bighead.Next
	bigtail.Next = nil
	return smallhead.Next
}

func main() {
	tail1 := CreateNewList()
	list2 := partition(tail1, 3)
	for list2 != nil {
		fmt.Println(list2.Val)
		list2 = list2.Next
	}
}
