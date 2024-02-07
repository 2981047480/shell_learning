package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func FindNumber(arr []int, target int) (index int) {
	var left, right, middle int
	left = 0
	right = len(arr) - 1
	for left <= right {
		middle = left + (right-left)>>1
		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] == target {
			return middle
		} else {
			left = middle + 1
		}
	}
	return -1
}

func FindNumberEasy(arr []int, target int) (index int) {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func RandomArray(len int) []int {
	arr := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr = append(arr, rand.Int())
	}
	return arr
}

func Validator() bool {
	rand.Seed(time.Now().UnixNano())
	len := 1000
	index := rand.Int() % len
	arr := RandomArray(len)
	sort.IntSlice.Sort(arr)
	if FindNumber(arr, arr[index]) != FindNumberEasy(arr, arr[index]) {
		return false
	} else {
		return true
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		if Validator() == false {
			panic("验证失败")
		}
	}
	fmt.Println("验证通过")
}
