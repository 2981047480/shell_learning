package main

import (
	"fmt"
)

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
	// arr := make([]int, 10)
	// var text string
	// fmt.Scanf("%s", &text)
	// arr := strings.Split(text, " ")
	// fmt.Println(arr)
	arr := []int{4, 3, 2, 1}
	arr = InsertSort(arr)
	fmt.Println(arr)
}
