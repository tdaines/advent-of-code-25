package day05_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day05"
)

func TestParseRange(t *testing.T) {
	var r day05.Range = day05.ParseRange("14374432572227-15942816853763")

	assert.Equal(t, 14374432572227, r.Start)
	assert.Equal(t, 15942816853763, r.End)
}

func TestRangeContains(t *testing.T) {
	var r day05.Range = day05.ParseRange("3-5")

	assert.Equal(t, false, r.Contains(1))
	assert.Equal(t, true, r.Contains(3))
	assert.Equal(t, true, r.Contains(4))
	assert.Equal(t, true, r.Contains(5))
	assert.Equal(t, false, r.Contains(8))
}

func TestParseIngredientsDB(t *testing.T) {
	var db = []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}

	var ranges, ingredients = day05.ParseIngredientsDB(db)

	assert.Equal(t, 3, ranges[0].Start)
	assert.Equal(t, 5, ranges[0].End)

	assert.Equal(t, 10, ranges[1].Start)
	assert.Equal(t, 20, ranges[1].End)

	assert.Equal(t, []int{1, 5, 8, 11, 17, 32}, ingredients)
}

func TestCountAvailableFreshIngredients(t *testing.T) {
	var db = []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}

	var ranges, ingredients = day05.ParseIngredientsDB(db)
	var count = day05.CountAvailableFreshIngredients(ranges, ingredients)

	assert.Equal(t, 3, count)
}

func TestCountAllFreshIngredients(t *testing.T) {
	var db = []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}

	var ranges, _ = day05.ParseIngredientsDB(db)
	var count = day05.CountAllFreshIngredients(ranges)

	assert.Equal(t, 14, count)
}
