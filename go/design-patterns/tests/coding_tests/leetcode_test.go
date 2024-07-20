package coding_tests

import (
	"design-patterns/src/coding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindArraysMedium(t *testing.T) {
	medium := coding.FindMedianSortedArrays(
		[]int{1, 3},
		[]int{2},
	)

	assert.Equal(t, medium, 2.000)

	medium = coding.FindMedianSortedArrays(
		[]int{1, 2},
		[]int{3, 4},
	)
	assert.Equal(t, medium, 2.50)

	medium = coding.FindMediumSortedArraysFaster(
		[]int{1, 3},
		[]int{2, 4, 5},
	)
	assert.Equal(t, medium, 3.0)
}
