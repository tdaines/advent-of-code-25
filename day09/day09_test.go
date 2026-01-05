package day09_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day09"
)

func TestFindLargestRect(t *testing.T) {
	var lines = []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}

	var area = day09.FindLargestRect(lines)
	assert.Equal(t, 50, area)
}
