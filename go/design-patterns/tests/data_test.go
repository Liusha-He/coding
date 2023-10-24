package tests

import (
	"design-patterns/src/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	df := data.ReadCSVFile("./data/cpi_from_1980.csv")
	assert.Equal(t, df.Names()[0], "DATE")
	l, w := df.Dims()
	assert.Equal(t, l, 518)
	assert.Equal(t, w, 2)
	assert.Equal(t, df.Names(), []string{"DATE", "CPIAUCSL"})
}
