package main

import (
	"fmt"
)

func FindPeak(arr []int) (index int) {
	var left, right, middle int
	left = 1
	right = len(arr) - 2 //注意这里的left和right的值，防止下面越界
	index = -1
	if len(arr) == 1 {
		return 0
	}
	if len(arr) == 2 && arr[0] == arr[1] {
		return index
	}
	if arr[0] > arr[1] {
		return 0
	}
	if arr[len(arr)-1] > arr[len(arr)-2] {
		return len(arr) - 1
	}
	for left <= right {
		middle = left + (right-left)>>1
		if arr[middle] < arr[middle-1] {
			right = middle - 1
			index = middle
		} else if arr[middle] < arr[middle+1] {
			left = middle + 1
			index = middle
		} else {
			return middle
		}
		// var leftval = -2 ^ 31
		// var rightval = -2 ^ 31
		// if middle == 0 {
		// 	leftval = -2 ^ 31
		// } else {
		// 	leftval = arr[middle-1]
		// }
		// if middle == len(arr)-1 {
		// 	rightval = -2 ^ 31
		// } else {
		// 	rightval = arr[middle+1]
		// }
		// if arr[middle] > leftval && arr[middle] > rightval {
		// 	return middle
		// } else if arr[middle] < leftval {
		// 	right = middle - 1
		// 	index = middle
		// } else if arr[middle] < rightval {
		// 	left = middle + 1
		// 	index = middle
		// }
	}
	// fmt.Println(arr, index)
	return index
}

func main() {
	arr := []int{3, 4, 3, 2, 1}
	fmt.Println(FindPeak(arr))
}
