package main

type Queue struct {
	arr  []int
	head int
	// tail int
	cap  int
	size int
}

func Init(cap int) Queue {
	arr := make([]int, cap)
	return Queue{
		arr: arr,
		cap: cap,
	}
}

func (q *Queue) Cap() int { return q.cap }
func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Pop() bool {
	if q.Isempty() {
		return false
	} else {
		q.head = (q.head + 1) % q.cap
		q.size--
		return true
	}
}

func (q *Queue) Isempty() bool {
	return q.size == 0
	// if q.size == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
	// if else去掉 节约了很多时间
}

func (q *Queue) Isfull() bool {
	return q.size == q.cap
	// if q.size == q.cap {
	// 	return true
	// } else {
	// 	return false
	// }
	// if else去掉 节约了很多时间
}

func (q *Queue) Push(item int) bool {
	if q.Isfull() {
		return false
	} else {
		index := (q.head + q.size) % q.cap
		q.arr[index] = item
		q.size++
		return true
	}
}

func (q *Queue) Front() int {
	if q.Isempty() {
		return -1
	} else {
		return q.arr[q.head]
	}
}

func (q *Queue) Rear() int {
	if q.Isempty() {
		return -1
	} else {
		index := (q.head + q.size - 1) % q.cap
		return q.arr[index]
		// if q.tail != 0 { // 这里需要考虑tail为0的情况
		// 	return q.arr[q.tail-1]
		// } else {
		// 	return q.arr[q.cap-1]
		// }
	}
}

func main() {

}
