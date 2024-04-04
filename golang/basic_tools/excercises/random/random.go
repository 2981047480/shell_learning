package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	flag = 1
)

func main() {
	add_int := 0
	plus_int := 1
	for i := 0; i < 20; i++ {
		item := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(20) + 1
		if item&flag == 0 {
			plus_int *= item
		} else {
			add_int += item
		}
	}
	fmt.Println("加法:", add_int, "乘法:", plus_int)
}
