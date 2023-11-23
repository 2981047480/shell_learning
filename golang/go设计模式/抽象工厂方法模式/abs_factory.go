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

type CreateChinaApple interface {
	CreateApple()
}

type CreateChinaBanana interface {
	CreateBanana()
}

type CreateAmiricaApple interface {
	CreateApple()
}

type CreateAmericaBanana interface {
	CreateBanana()
}

type Apple struct {
	Fruit
}
type Banana struct {
	Fruit
}

type ChinaApple struct {
	AbsApple
}
type ChinaBanana struct {
	AbsBanana
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("I'm a China apple")
}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("I'm a China banana")
}

type AmeracaApple struct {
	AbsApple
}

type AmeracaBanana struct {
	AbsBanana
}

func (aa *AmeracaApple) ShowApple() {
	fmt.Println("I'm a Ameraca apple")
}

func (ab *AmeracaBanana) ShowBanana() {
	fmt.Println("I'm a Ameraca banana")
}

func (ca *CreateChinaApple) CreateApple() (ca *Apple) {
	ca = new(Apple)
	ca.Fruit = new(ChinaApple)
	return ca
}

func main() {
	ca = new(ChinaApple)
}
