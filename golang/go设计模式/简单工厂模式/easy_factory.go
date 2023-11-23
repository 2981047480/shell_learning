package main

import (
	"fmt"
)

type Fruit interface {
	String() string
}
type Apple struct {
	f Fruit
}

func (apple *Apple) String() string {
	// fmt.Println("I am apple")
	return "I am apple"
}

type Banana struct {
	f Fruit
}

func (banana *Banana) String() string {
	// fmt.Println("I am banana")
	return "I am banana"
}

type Factory struct{}

func (factory *Factory) CreateFactory(str string) Fruit {
	var fruit Fruit
	switch str {
	case "apple":
		fruit = new(Apple)
	case "banana":
		fruit = new(Banana)
	default:
		fmt.Println("Unknown fruit")
	}
	return fruit
}

func main() {
	apple := Apple{}
	banana := Banana{}
	fmt.Println(apple.String(), banana.String())
	// apple.String()
	// banana.String()
	factory := new(Factory)
	apple2 := factory.CreateFactory("apple")
	// apple2.String()
	fmt.Println(apple2.String())
}
