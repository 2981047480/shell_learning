package main

import "fmt"

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
func (s *Stack) Peek() int {
	if s.Isempty() {
		return -1
	}
	item := s.arr[s.size-1]
	return item
}

type MinStack struct {
	datastack Stack
	minstack  Stack
}

func Constructor() MinStack {
	return MinStack{datastack: Stack{arr: make([]int, 8001), cap: 8001}, minstack: Stack{arr: make([]int, 8001), cap: 8001}}
}

func (this *MinStack) Push(val int) {
	this.datastack.Push(val)
	if this.minstack.Isempty() {
		this.minstack.Push(val)
	} else {
		if this.minstack.Peek() > val {
			this.minstack.Push(val)
		} else {
			this.minstack.Push(this.minstack.Peek())
		}
	}
}

func (this *MinStack) Pop() {
	this.datastack.Pop()
	this.minstack.Pop()
	// fmt.Printf("data %v\n", this.datastack.Pop())
	// fmt.Printf("min %v\n", this.minstack.Pop())
}

func (this *MinStack) Top() int {
	val := this.datastack.Peek()
	this.minstack.Peek()
	return val
}

func (this *MinStack) GetMin() int {
	return this.minstack.Peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	for i := 0; i < obj.datastack.size; i++ {
		fmt.Println(obj.datastack.Pop())
		fmt.Println(obj.minstack.Pop())
	}
}
