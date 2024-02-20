package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(list *ListNode, x int) *ListNode {
	if list == nil {
		return list
	}
	var head1 = new(ListNode)
	var list1 = new(ListNode)
	list1.Next = nil
	head1 = list1
	var head2 = new(ListNode)
	var list2 = new(ListNode)
	list2.Next = nil
	head2 = list2
	var next = new(ListNode)
	for list != nil {
		if list.Val < x {
			next = list
			list = list.Next
			next.Next = list1.Next
			list1.Next = next
			list1 = list1.Next
		} else {
			next = list
			list = list.Next
			next.Next = list2.Next
			list2.Next = next
			list2 = list2.Next
		}
	}
	list1.Next = head2.Next
	return head1.Next
}

func main() {

}
