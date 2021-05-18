package main

import "fmt"

// 定义一个女孩的接口让下面的类隐式继承
type gril interface {
	weight()
}

type FatGril struct{}

func (FatGril) weight() {
	fmt.Println("100kg")
}

type ThinGirl struct{}

func (ThinGirl) weight() {
	fmt.Println("45kg")
}

// GirlFactory 简单工厂
type GirlFactory struct{}

func (*GirlFactory) CreateGirl(like string) gril {
	switch like {
	case "fat":
		return &FatGril{}
	case "thin":
		return &ThinGirl{}
	default:
		return nil
	}
}

func main() {
	// 测试
	factor := &GirlFactory{}

	var g gril

	g = factor.CreateGirl("fat")
	g.weight()
	g = factor.CreateGirl("thin")
	g.weight()
}
