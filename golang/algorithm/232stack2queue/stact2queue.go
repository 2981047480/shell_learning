package main

type Stack struct {
	arr  []int
	size int
	cap  int
}

func (s *Stack) Size() int { return s.size }
func (s *Stack) Isempty() bool {
	return s.size == 0
}
func (s *Stack) Isfull() bool {
	return s.size == s.cap
}
func (s *Stack) Push(item int) bool {
	if s.Isfull() {
		return false
	}
	s.arr[s.size] = item
	s.size++
	return true
}
func (s *Stack) Pop() int {
	if s.Isempty() {
		return -1
	}
	item := s.arr[s.size-1]
	s.size--
	return item
}

type MyQueue struct {
	stackA Stack
	stackB Stack
}

func Constructor() MyQueue {
	return MyQueue{stackA: Stack{arr: make([]int, 100), cap: 100}, stackB: Stack{arr: make([]int, 100), cap: 100}}
}

func (this *MyQueue) atob() {
	for this.stackA.Isempty() == false {
		this.stackB.Push(this.stackA.Pop())
	}
}

func (this *MyQueue) Push(x int) bool {
	return this.stackA.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stackB.Isempty() == false {
		return this.stackB.Pop()
	} else {
		this.atob()
		return this.stackB.Pop()
	}
}

func (this *MyQueue) Peek() int {
	if this.stackB.Isempty() == false {
		return this.stackB.arr[this.stackB.size-1]
	} else {
		this.atob()
		return this.stackB.arr[this.stackB.size-1]
	}
}

func (this *MyQueue) Empty() bool {
	return this.stackA.Isempty() == true && this.stackB.Isempty() == true
	// if this.stackA.Isempty() == true && this.stackB.Isempty() == true {
	// 	return true
	// } else {
	// 	return false
	// }
	// 这个记住了！！！！！！！因为这个差了1ms
}
