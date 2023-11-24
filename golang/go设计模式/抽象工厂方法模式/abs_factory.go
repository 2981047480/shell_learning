package main

import (
	"fmt"
)

type Fruit interface{}
type AbsApple interface {
	ShowApple()
}
type AbsBanana interface {
	ShowBanana()
}

type Apple struct{}
type Banana struct{}

func (a *Apple) ShowApple() {
	fmt.Println("I am an apple")
}
func (b *Banana) ShowBanana() {
	fmt.Println("I am a banana")
}

type CreateFruitFactory interface {
	CreateApple() *AbsApple
	CreateBanana() *AbsBanana
}

type AmericaFactory struct{}

func (aa *AmericaFactory) CreateApple() *Apple {
	return new(Apple)
}

func (ab *AmericaFactory) CreateBanana() *Banana {
	return new(Banana)
}

type ChinaFactory struct{}

func (cb *ChinaFactory) CreateBanana() *Banana {
	return new(Banana)
}

func (ca *ChinaFactory) CreateApple() *Apple {
	return new(Apple)
}

func main() {
	af := new(AmericaFactory)
	aa := af.CreateApple()
	aa.ShowApple()
	ab := af.CreateBanana()
	ab.ShowBanana()
	cf := new(ChinaFactory)
	ca := cf.CreateApple()
	ca.ShowApple()
	cb := cf.CreateBanana()
	cb.ShowBanana()
}
