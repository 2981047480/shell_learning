package list

// type node struct {
// 	Val int
// }

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
