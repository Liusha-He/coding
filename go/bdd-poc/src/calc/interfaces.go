package calc

import "fmt"

type ICalculator interface {
	SayHello() string
	AddTwo(int) int
}

type MyCalculator struct {
	Name string
}

func (mc MyCalculator) SayHello() string {
	return fmt.Sprintf("hello %s", mc.Name)
}

func (mc MyCalculator) AddTwo(first int) int {
	return first + 2
}

func NewCalculator(name string) *MyCalculator {
	return &MyCalculator{
		Name: name,
	}
}
