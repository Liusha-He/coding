package construction_tests

import (
	"strings"
	"testing"

	"design-patterns/src/construction_patterns"
	assert "github.com/stretchr/testify/assert"
)

func TestStringBuilder(t *testing.T) {
	hello := "hello"

	builder := strings.Builder{}

	builder.WriteString("<p>")
	builder.WriteString(hello)
	builder.WriteString("</p>")

	assert.Equal(t, builder.String(), "<p>hello</p>",
		"string builder output incorrect.")
}

func TestFacets(t *testing.T) {
	b := construction_patterns.NewPersonBuilder()
	b.AddAddress("test street", "123 456", "London")
	b.AddOccupationInfo("ki", "developer", 61000.00)

	assert.Equal(t, b.Person.String(),
		"test street/123 456/London",
	)
}

func TestParameters(t *testing.T) {
	// do something here.
}

func TestFunctional(t *testing.T) {
	// do something here...
}
