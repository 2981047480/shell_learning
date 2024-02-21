package main

type Queue struct {
	arr  []int
	head int
	size int
}

func (this *Queue) Push(value int) bool {
	if this.IsFull() {
		return false
	}

	index := (this.head + this.size) % len(this.arr)
	this.arr[index] = value
	this.size++

	return true
}

func (this *Queue) Pop() int {
	if this.IsEmpty() {
		return -1
	}

	item := this.arr[this.head]
	this.head = (this.head + 1) % len(this.arr)
	this.size--

	return item
}

func (this *Queue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.arr[this.head]
	}
}

func (this *Queue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.arr[(this.head+this.size-1)%len(this.arr)]
	}
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) IsFull() bool {
	return this.size == len(this.arr)
}

type MyStack struct {
	Que Queue
}

func Constructor() MyStack {
	return MyStack{Que: Queue{arr: make([]int, 100)}}
}

func (this *MyStack) Push(x int) {
	this.Que.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.Que.size-1; i++ {
		this.Que.Push(this.Que.Pop())
	}
	return this.Que.Pop()
}

func (this *MyStack) Top() int {
	for i := 0; i < this.Que.size-1; i++ {
		this.Que.Push(this.Que.Pop())
	}
	item := this.Que.Pop()
	this.Que.Push(item)
	return item
}

func (this *MyStack) Empty() bool {
	return this.Que.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
