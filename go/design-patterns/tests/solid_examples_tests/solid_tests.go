package solid_examples_tests

import (
	"fmt"
	"os"
	"testing"

	"design-patterns/src/solid_examples"

	"github.com/stretchr/testify/assert"
)

func TestSimpleResponsibility(t *testing.T) {
	j := solid_examples.Journal{}

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

func TestOpenClosed(t *testing.T) {
	// old way
	apple := solid_examples.Product{"apple", solid_examples.Green, solid_examples.Small}
	tree := solid_examples.Product{"tree", solid_examples.Green, solid_examples.Large}
	house := solid_examples.Product{"house", solid_examples.Blue, solid_examples.Large}

	products := []solid_examples.Product{apple, tree, house}

	f := solid_examples.Filter{}
	expectedNames := []string{"apple", "tree"}

	for i, p := range f.FilterByColor(products, solid_examples.Green) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	// specification
	fmt.Println("use the new method")

	bf := solid_examples.BetterFilter{}
	greenspec := solid_examples.ColorSprcification{solid_examples.Green}

	for i, p := range bf.Filter(products, greenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largespec := solid_examples.SizeSpecification{solid_examples.Large}

	expectedNames = []string{"tree", "house"}

	for i, p := range bf.Filter(products, largespec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largegreenspec := solid_examples.AndSpecification{greenspec, largespec}
	expectedNames = []string{"tree"}

	for i, p := range bf.Filter(products, largegreenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect")
	}
}

func TestLiskovSubstitution(t *testing.T) {
	rc := &solid_examples.Rectangle{Width: 2, Height: 3}
	ex, ac := solid_examples.UsedSized(rc)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))

	sq := solid_examples.Square{Size: 5}.Rectangle()
	ex, ac = solid_examples.UsedSized(&sq)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))
}

func TestInterfaceSegregation(t *testing.T) {
	// no test here
}

func TestDependencyInversion(t *testing.T) {
	parent := solid_examples.Person{"John"}
	child1 := solid_examples.Person{"Chris"}
	child2 := solid_examples.Person{"Finlay"}

	relations := solid_examples.Relationships{}

	relations.AddChildAndParent(&parent, &child1)
	relations.AddChildAndParent(&parent, &child2)

	r := solid_examples.Research{&relations}

	foundNames := r.Investigate("John")
	assert.Equal(t, foundNames, "Chris,Finlay", "found names wrong.")
}
