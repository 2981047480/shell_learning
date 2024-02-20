package main

import (
	"fmt"

	"example.com/m/v2/algorithm/list"
)

func merge(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	list_head := new(list.ListNode)
	list_head.Next = nil
	merge_list := list_head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			merge_list.Next = list2
			break
		}
		if list2 == nil {
			merge_list.Next = list1
			break
		}
		fmt.Println(list1, list2)
		if list1.Val < list2.Val {
			merge_list.Next = list1
			list1 = list1.Next
			merge_list = merge_list.Next
		} else {
			merge_list.Next = list2
			list2 = list2.Next
			merge_list = merge_list.Next
		}
	}
	return list_head.Next
}

func main() {
	l1 := list.CreateNewList()
	l2 := list.CreateNewList()
	merge(l1, l2)
}
