package day07_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day07"
)

func TestTraverseManifold(t *testing.T) {
	var manifold = [][]byte{
		[]byte(".......S......."),
		[]byte("..............."),
		[]byte(".......^......."),
		[]byte("..............."),
		[]byte("......^.^......"),
		[]byte("..............."),
		[]byte(".....^.^.^....."),
		[]byte("..............."),
		[]byte("....^.^...^...."),
		[]byte("..............."),
		[]byte("...^.^...^.^..."),
		[]byte("..............."),
		[]byte("..^...^.....^.."),
		[]byte("..............."),
		[]byte(".^.^.^.^.^...^."),
		[]byte("..............."),
	}

	var numSplits = day07.TraverseManifold(manifold)

	assert.Equal(t, 21, numSplits)
}

func TestTraverseQuantumManifold(t *testing.T) {
	var manifold = [][]byte{
		[]byte(".......S......."),
		[]byte("..............."),
		[]byte(".......^......."),
		[]byte("..............."),
		[]byte("......^.^......"),
		[]byte("..............."),
		[]byte(".....^.^.^....."),
		[]byte("..............."),
		[]byte("....^.^...^...."),
		[]byte("..............."),
		[]byte("...^.^...^.^..."),
		[]byte("..............."),
		[]byte("..^...^.....^.."),
		[]byte("..............."),
		[]byte(".^.^.^.^.^...^."),
		[]byte("..............."),
	}

	var numSplits = day07.TraverseManifold(manifold)

	assert.Equal(t, 21, numSplits)

	var numPaths = day07.TraverseQuantumManifold(manifold)

	assert.Equal(t, 40, numPaths)
}
