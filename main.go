package main

import (
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/tdaines/advent-of-code-25/day01"
)

func main() {
	var dayBreak = "---"
	var partBreak = "----"
	var answerBreak = "-------------"
	var elapsedBreak = "-----------"

	t := tabby.New()
	t.AddHeader("DAY", "PART", "ANSWER", "ELAPSED")
	var total time.Duration = 0

	var answer, elapsed = day01.Part1()
	total += elapsed
	t.AddLine("1", "1", answer, elapsed)

	// answer, elapsed = day01.Part2()
	// total += elapsed
	// t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	t.AddLine("", "", "TOTAL", total)
	t.Print()
}
