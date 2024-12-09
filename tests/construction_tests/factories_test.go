package construction_tests

import (
	cp "design-patterns/src/construction_patterns"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactories(t *testing.T) {
	factory1 := cp.NewEmployFactory("Developer", 80000.00)
	emp1 := factory1("Liusha He")

	factory2 := cp.EmployFactory{Position: "Manager", AnnualIncome: 100000.00}
	emp2 := factory2.Create("Degere")

	assert.Equal(t, emp1.SayHello(),
		"Hello my name is  and my position is Developer, my salary is 80000.00")
	assert.Equal(t, emp2.SayHello(),
		"Hello my name is Degere and my position is Manager, my salary is 100000.00")
}
