package main

import (
	"fmt"
)

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() (fruit *Fruit)
}

type Apple struct {
	Fruit
}

func (a *Apple) Show() {
	fmt.Println("I am an apple")
}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {
	fmt.Println("I am a banana")
}

type AppleFactory struct {
	AbstractFactory
}

type BananaFactory struct {
	AbstractFactory
}

func (af *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
	// var fruit Fruit
	// fruit = new(Apple)
	// return fruit
}

func (bf *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
	// var fruit Fruit
	// fruit = new(Banana)
	// return fruit
}

func main() {
	af := new(AppleFactory).CreateFruit()
	bf := new(BananaFactory).CreateFruit()
	af.Show()
	bf.Show()
}
