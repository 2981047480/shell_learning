package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func RandomArray(len int, max int) []int {
	arr := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr = append(arr, rand.Int()%max)
	}
	return arr
}

func Validator() bool {
	rand.Seed(time.Now().UnixNano())
	len := 10
	max := 5
	index := rand.Int() % len
	arr := RandomArray(len, max)
	sort.IntSlice.Sort(arr)
	if FindLeft(arr, arr[index]) != FindLeftEasy(arr, arr[index]) {
		return false
	} else {
		return true
	}
}

func FindLeft(arr []int, target int) (index int) {
	var left, right, middle int
	left = 0
	right = len(arr) - 1
	index = -1
	for left <= right {
		middle = left + (right-left)>>1
		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] == target {
			index = middle
			left = middle + 1
		} else {
			left = middle + 1
		}
	}
	fmt.Println(arr, target, index)
	return index
}

func FindLeftEasy(arr []int, target int) (index int) {
	index = -1
	for i, v := range arr {
		if v == target {
			index = i
		}
	}
	return index
}

func main() {
	for i := 0; i < 10; i++ {
		if Validator() == false {
			panic("验证失败")
		}
	}
	fmt.Println("验证通过")
}
