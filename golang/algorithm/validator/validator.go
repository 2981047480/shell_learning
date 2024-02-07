package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Validator() {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(20)
	}
	arr1 := make([]int, 100)
	copy(arr1, arr)
	arr2 := make([]int, 100)
	copy(arr2, arr)
	arr3 := make([]int, 100)
	copy(arr3, arr)
	arr = BubleSort(arr)
	arr1 = SelectSort(arr1)
	arr2 = InsertSort(arr2)
	arr3 = InsertSort(arr3)
	if IntSliceEqual(arr, arr2) && IntSliceEqual(arr, arr1) && IntSliceEqual(arr, arr3) {
		fmt.Println("验证成功")
	} else {
		fmt.Println("验证失败")
	}
}

func IntSliceEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func BubleSort(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	for i := len(arr); i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func BubleSort1(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	j := 0
	for i := len(arr); i >= 0; i-- {
		// for j := 0; j < i-1; j++ {
		if arr[j] > arr[j+1] {
			swap(arr, j, j+1)
		}
		if j == i-1 {
			j = 0
		} else {
			j++
		}
		// }
	}
	return arr
}

func SelectSort(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		minindex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		swap(arr, i, minindex)
	}
	return arr
}

func InsertSort(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			} else {
				break
			}
		}
	}
	return arr
}

func swap(arr []int, index1 int, index2 int) {
	arr[index1], arr[index2] = arr[index2], arr[index1]
}

func main() {
	Validator()
}
