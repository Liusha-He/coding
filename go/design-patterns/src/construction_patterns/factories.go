package construction_patterns

import "fmt"

type Employee interface {
	SayHello() string
}

type employee struct {
	Name, Position string
	AnnualIncome   float32
}

func (e employee) SayHello() string {
	return fmt.Sprintf("Hello my name is %s and my position is %s, my salary is %.2f",
		e.Name,
		e.Position,
		e.AnnualIncome)
}

// functional
func NewEmployFactory(position string,
	annualIncome float32) func(name string) Employee {
	return func(name string) Employee {
		return employee{Position: position,
			AnnualIncome: annualIncome}
	}
}

// structural approach
type EmployFactory struct {
	Position     string
	AnnualIncome float32
}

func (f EmployFactory) Create(name string) Employee {
	return employee{
		Position:     f.Position,
		AnnualIncome: f.AnnualIncome,
		Name:         name,
	}
}
