package main

type Queue struct {
	arr  []int
	head int
	tail int
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
		// q.head++
		// if q.head == q.cap {
		// 	q.head = 0
		// }
		q.head = (q.head + 1) % q.cap
		q.size--
		return true
	}
}

func (q *Queue) Isempty() bool {
	if q.size == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Isfull() bool {
	if q.size == q.cap {
		return true
	} else {
		return false
	}
}

func (q *Queue) Push(item int) bool {
	if q.Isfull() {
		return false
	} else {
		q.arr[q.tail] = item
		q.tail = (q.tail + 1) % q.cap
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
		if q.tail != 0 { // 这里需要考虑tail为0的情况
			return q.arr[q.tail-1]
		} else {
			return q.arr[q.cap-1]
		}
	}
}

func main() {

}
