package main

import (
	"fmt"
	"os"
)

type Args struct {
	args []string
}

func (a *Args) Read_Args() []string {
	var str []string
	for _, arg := range os.Args[1:] {
		// fmt.Printf(" %s", arg)
		str = append(str, arg)
	}
	a.args = str
	return str
}

func main() {
	// var str []string
	// for _, arg := range os.Args[1:] {
	// 	// fmt.Printf(" %s", arg)
	// 	str = append(str, arg)
	// }
	a := new(Args)
	a.Read_Args()
	fmt.Println("Args ", a.args)
	// fmt.Println(str)
}
