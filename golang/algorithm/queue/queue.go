package main

type Queue struct {
	arr  []int
	head int
	tail int
	cap  int
	size int
}

func Constructor(k int) Queue {
	arr := make([]int, k)
	return Queue{
		arr: arr,
		cap: k,
	}
}

func (this *Queue) Push(value int) bool {
	if this.IsFull() {
		return false
	} else {
		index := (this.head + this.size) % this.cap
		this.arr[index] = value
		this.size++
		return true
	}
}

func (this *Queue) Pop() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.head = (this.head + 1) % this.cap
		this.size--
		return true
	}
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
		index := (this.head + this.size - 1) % this.cap
		return this.arr[index]
		// if q.tail != 0 { // 这里需要考虑tail为0的情况
		// 	return q.arr[q.tail-1]
		// } else {
		// 	return q.arr[q.cap-1]
		// }
	}
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
	// if this.size == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
}

func (this *Queue) IsFull() bool {
	return this.size == this.cap
	// if this.size == this.cap{
	//     return true
	// }else{
	//     return false
	// }
}

/**
 * Your Queue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
