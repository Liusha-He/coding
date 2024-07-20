package syntax_sugar_tests

import (
	"fmt"
	"os"
	"testing"

	"design-patterns/src/syntax_sugar"

	"github.com/stretchr/testify/assert"
)

func TestUskovSubstitution(t *testing.T) {
	rc := &syntax_sugar.Rectangle{Width: 2, Height: 3}
	ex, ac := syntax_sugar.UsedSized(rc)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))

	sq := syntax_sugar.Square{Size: 5}.Rectangle()
	ex, ac = syntax_sugar.UsedSized(&sq)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))
}

func TestDependencyInversion(t *testing.T) {
	parent := syntax_sugar.Person{"John"}
	child1 := syntax_sugar.Person{"Chris"}
	child2 := syntax_sugar.Person{"Finlay"}

	relations := syntax_sugar.Relationships{}

	relations.AddChildAndParent(&parent, &child1)
	relations.AddChildAndParent(&parent, &child2)

	r := syntax_sugar.Research{&relations}

	foundNames := r.Investigate("John")
	assert.Equal(t, foundNames, "Chris,Finlay", "found names wrong.")
}

func TestOpenClosed(t *testing.T) {
	// old way
	apple := syntax_sugar.Product{"apple", syntax_sugar.Green, syntax_sugar.Small}
	tree := syntax_sugar.Product{"tree", syntax_sugar.Green, syntax_sugar.Large}
	house := syntax_sugar.Product{"house", syntax_sugar.Blue, syntax_sugar.Large}

	products := []syntax_sugar.Product{apple, tree, house}

	f := syntax_sugar.Filter{}
	expectedNames := []string{"apple", "tree"}

	for i, p := range f.FilterByColor(products, syntax_sugar.Green) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	// specification
	fmt.Println("use the new method")

	bf := syntax_sugar.BetterFilter{}
	greenspec := syntax_sugar.ColorSprcification{syntax_sugar.Green}

	for i, p := range bf.Filter(products, greenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largespec := syntax_sugar.SizeSpecification{syntax_sugar.Large}

	expectedNames = []string{"tree", "house"}

	for i, p := range bf.Filter(products, largespec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largegreenspec := syntax_sugar.AndSpecification{greenspec, largespec}
	expectedNames = []string{"tree"}

	for i, p := range bf.Filter(products, largegreenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect")
	}
}

func TestSimpleResponse(t *testing.T) {
	j := syntax_sugar.Journal{}

	j.AddEntry("hello")
	j.AddEntry("world")
	j.AddEntry("how")
	j.AddEntry("are")
	j.AddEntry("you")

	j.Save("log.txt")

	j.Load("log.txt")

	assert.Equal(
		t,
		"1: hello\n2: world\n3: how\n4: are\n5: you\n1: hello\n2: world\n3: how\n4: are\n5: you",
		j.String(),
		"output incorrect",
	)

	j.RemoveEntry(3)

	assert.Equal(
		t,
		"1: hello\n2: world\n3: how\n5: you\n1: hello\n2: world\n3: how\n4: are\n5: you",
		j.String(),
		"output incorrect...",
	)

	os.Remove("log.txt")

}

func TestInterfaceSegregation(t *testing.T) {
	// no test here
}
