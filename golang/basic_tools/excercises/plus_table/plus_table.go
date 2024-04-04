package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j > i {
				continue
			}
			// 方法1
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
			// 方法2
			var tab string
			if i*j < 10 {
				tab = "  "
			} else {
				tab = " "
			}
			fmt.Printf("%d*%d=%d%s", j, i, i*j, tab)
		}
		fmt.Printf("\n")
	}
}
