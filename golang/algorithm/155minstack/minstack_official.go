package minstack_quickly

type MinStack struct {
	data [][2]int
}

func Constructor() MinStack {
	return MinStack{
		data: make([][2]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	minVal := val

	if length := len(this.data); length > 0 {
		if this.data[length-1][1] < minVal {
			minVal = this.data[length-1][1]
		}
	}

	this.data = append(this.data, [2]int{val, minVal})
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
