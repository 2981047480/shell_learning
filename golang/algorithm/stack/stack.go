package stack

import "fmt"

type Stack struct {
	arr  []int
	size int
	cap  int
}

func (s *Stack) Len() int { return s.size }
func (s *Stack) Cap() int { return s.cap }
func (s *Stack) Init(n int) {
	s.arr = make([]int, n)
	s.cap = n
	s.size = 0
}
func (s *Stack) Isempty() bool {
	if s.size == 0 {
		return true
	} else {
		return false
	}
}
func (s *Stack) Push(item int) {
	s.arr[s.size] = item
	s.size++
}
func (s *Stack) Pop() int {
	top := s.arr[s.size-1]
	s.size--
	return top
}
func (s *Stack) Peak() int { return s.arr[s.size-1] }

func test() {
	s := Stack{}
	s.Init(10)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peak())
	fmt.Printf("%v %v %v %v\n", s.arr, s.Peak(), s.Len(), s.Cap())
	fmt.Println(s.Pop())
	fmt.Printf("%v %v %v %v\n", s.arr, s.Peak(), s.Len(), s.Cap())
	s.Push(3)
	fmt.Printf("%v %v %v %v\n", s.arr, s.Peak(), s.Len(), s.Cap())
}
