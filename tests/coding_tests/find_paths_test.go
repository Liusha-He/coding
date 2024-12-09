package coding_tests

import (
	"design-patterns/src/coding"
	"testing"
)

func TestFindPaths(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 1, 1},
	}

	n := coding.FindPaths(grid)

	if n != 4 {
		t.Errorf("Expected 4 paths, but get %d", n)
	}

	grid = [][]int{
		{1, 1, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}
	n = coding.FindPaths(grid)
	if n != 5 {
		t.Errorf("Expected 6 paths, but get %d", n)
	}
}
