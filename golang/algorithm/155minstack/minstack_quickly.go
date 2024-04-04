package minstack_quickly

type MinStack struct {
	datastack []int
	minstack  []int
	size      int
	cap       int
}

func Constructor() MinStack {
	return MinStack{datastack: make([]int, 8001), minstack: make([]int, 8001), cap: 8001}
}

func (this *MinStack) Push(val int) {
	if this.size == this.cap {
		return
	}
	this.datastack[this.size] = val
	if this.size == 0 {
		this.minstack[this.size] = val
	} else {
		if this.minstack[this.size-1] > val {
			this.minstack[this.size] = val
		} else {
			this.minstack[this.size] = this.minstack[this.size-1]
		}
	}
	this.size++
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}
	this.size--
}

func (this *MinStack) Top() int {
	return this.datastack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
