package main

//这个题和最优解有差异 大概在7ms 最优解是0ms
import (
	"fmt"

	"example.com/m/v2/algorithm/list"
)

func add(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	var cover int
	head := new(list.ListNode)
	current := head
	for list1 != nil || list2 != nil {
		var list1_val int
		var list2_val int
		if list1 == nil {
			list1_val = 0
		} else {
			list1_val = list1.Val
		}
		if list2 == nil {
			list2_val = 0
		} else {
			list2_val = list2.Val
		}
		current.Val = (list1_val + list2_val + cover) % 10
		cover = (list1_val + list2_val + cover) / 10
		if list1 != nil {
			list1 = list1.Next
		}
		if list2 != nil {
			list2 = list2.Next
		}
		if list1 != nil || list2 != nil {
			current.Next = new(list.ListNode)
			current = current.Next
		}
	}
	if cover == 1 {
		current.Next = new(list.ListNode)
		current = current.Next
		current.Val = 1
		current.Next = nil
	}
	return head
}

func main() {
	l1 := list.CreateNewList()
	l2 := list.CreateNewList()
	l3 := add(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
