package regulartests

import (
	"design-patterns/src/calc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInherit(t *testing.T) {
	myCalc := calc.NewCalculator("calc")

	assert.Implements(t, (*calc.ICalculator)(nil), myCalc)
	assert.Equal(t, myCalc.SayHello(), "hello calc")
	assert.Equal(t, myCalc.AddTwo(10), 12)
}
