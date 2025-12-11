package day05

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")

	var ranges, ingredients = ParseIngredientsDB(lines)
	answer = CountAvailableFreshIngredients(ranges, ingredients)

	return answer, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")

	var ranges, _ = ParseIngredientsDB(lines)
	answer = CountAllFreshIngredients(ranges)

	return answer, time.Since(start)
}

type Range struct {
	Start int
	End   int
}

func ParseRange(line string) (r Range) {
	// 14374432572227-15942816853763
	var parts = strings.Split(line, "-")

	r.Start, _ = strconv.Atoi(parts[0])
	r.End, _ = strconv.Atoi(parts[1])

	return
}

func (r Range) Contains(value int) bool {
	return value >= r.Start && value <= r.End
}

func (r Range) Overlaps(other Range) bool {
	if r.Contains(other.Start) {
		return true
	}

	if r.Contains(other.End) {
		return true
	}

	return false
}

func (r *Range) Union(other Range) {
	r.Start = min(r.Start, other.Start)
	r.End = max(r.End, other.End)
}

func (r Range) NumIngredients() int {
	return (r.End - r.Start) + 1
}

func ParseIngredientsDB(db []string) (ranges []Range, ingredientIDs []int) {
	// 3-5
	// 10-14
	// 16-20
	// 12-18
	//
	// 1
	// 5
	// 8
	// 11
	// 17
	// 32

	var i = 0

	// Parse Ranges
	for {
		if i >= len(db) {
			break
		}

		var line = db[i]
		if len(line) == 0 {
			break
		}

		ranges = append(ranges, ParseRange(line))
		i++
	}

	sortRanges(ranges)
	ranges = mergeRanges(ranges)

	// Skip blank line
	i++

	// Parse Ingredient IDs
	for {
		if i >= len(db) {
			break
		}

		var line = db[i]
		if len(line) == 0 {
			break
		}

		var id, _ = strconv.Atoi(line)
		ingredientIDs = append(ingredientIDs, id)

		i++
	}

	return
}

func sortRanges(ranges []Range) {
	sort.Slice(ranges, func(i, j int) bool {
		var left = ranges[i]
		var right = ranges[j]

		if left.Start < right.Start {
			return true
		}

		return false
	})
}

func mergeRanges(ranges []Range) (mergedRanges []Range) {
	mergedRanges = append(mergedRanges, ranges[0])
	var currentRange = &mergedRanges[len(mergedRanges)-1]

	for i := 1; i < len(ranges); i++ {
		if currentRange.Overlaps(ranges[i]) {
			currentRange.Union(ranges[i])
		} else {
			mergedRanges = append(mergedRanges, ranges[i])
			currentRange = &mergedRanges[len(mergedRanges)-1]
		}
	}

	return mergedRanges
}

func CountAvailableFreshIngredients(ranges []Range, ingredientIDs []int) (count int) {
	for _, id := range ingredientIDs {
		for _, r := range ranges {
			if r.Contains(id) {
				count++
				break
			}
		}
	}

	return
}

func CountAllFreshIngredients(ranges []Range) (sum int) {
	for _, r := range ranges {
		sum += r.NumIngredients()
	}

	return
}
