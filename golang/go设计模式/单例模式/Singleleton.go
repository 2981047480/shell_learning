package main

import (
	"fmt"
	"sync"
)

// type get_instance interface {
// 	Get_instance() *instance
// }
var once sync.Once

type instance struct {
	id          int
	value       int
	description string
}

func print_struct(ins *instance) {
	fmt.Println(ins.id, ins.value, ins.description)
}

func Get_instance() *instance {
	once.Do(func() {
		fmt.Println("get_instance")
	})
	return new(instance)
}

func main() {
	ins := Get_instance()
	// fmt.Println(ins)
	ins.id = 1
	ins.value = 10000
	ins.description = "test"
	print_struct(ins)
	ins2 := Get_instance()
	print_struct(ins2)
}
