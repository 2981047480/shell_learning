package main

import (
	"fmt"
)

type Item interface {
	// GetName() string
	GetPrice() float64
}

// 抽象商品类

type Proxy interface {
	Buy()
}

// 抽象代理

type CreateProxy interface {
	GetProxy() Proxy
}

// 用来创建代理的类

type Glasses struct {
	price float64
}

// 实体货物

func (g *Glasses) GetPrice() float64 {
	return g.price
}

// 上面商品类的实现

type Mary struct {
	proxy Proxy
	count int
}

func (m *Mary) Buy() {
	defer func(m *Mary) { m.count++ }(m)
	fmt.Println("Mary buys a glasses")
}

// 代购类的实现

func (m *Mary) GetProxy() *Proxy {
	return new(Proxy)
}

func main() {
	mary := Mary{count: 0}
	glasses := Glasses{price: 100.0}
	mary.Buy()
	fmt.Println(mary.count)
	mary.Buy()
	fmt.Println(mary.count)
	mary.Buy()
	fmt.Println(mary.count)
	fmt.Println(glasses.GetPrice())
}
