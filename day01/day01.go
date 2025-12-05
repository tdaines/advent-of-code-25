package day01

import (
	_ "embed"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")
	var password = GetPassword(lines)

	return password, time.Since(start)
}

type Direction rune

const (
	Right Direction = 'R'
	Left  Direction = 'L'
)

const (
	DIAL_MIN  = 0
	DIAL_MAX  = 99
	NUM_TICKS = (DIAL_MAX - DIAL_MIN) + 1
)

func GetPassword(lines []string) int {
	var dial = 50
	var password = 0

	for _, line := range lines {
		direction, distance := ParseRotation(line)
		dial = RotateDial(dial, direction, distance)
		if dial == 0 {
			password += 1
		}
	}

	return password
}

func ParseRotation(line string) (direction Direction, distance int) {
	if len(line) == 0 {
		return
	}

	if line[0] == 'R' {
		direction = Right
	} else {
		direction = Left
	}

	distance, err := strconv.Atoi(line[1:])
	if err == nil {
		return
	}

	distance = 0
	return
}

func RotateDial(dial int, direction Direction, distance int) int {
	distance = distance % NUM_TICKS

	if direction == Right {
		dial = (dial + distance) % NUM_TICKS
	} else {
		dial -= distance
		if dial < 0 {
			dial = NUM_TICKS + dial
		}
	}

	return dial
}
