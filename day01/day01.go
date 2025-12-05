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
	var password = GetPassword(lines, false)

	return password, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")
	var password = GetPassword(lines, true)

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

func GetPassword(lines []string, includeZeroTouches bool) int {
	var dial = 50
	var password = 0
	var zeroTouches = 0

	for _, line := range lines {
		direction, distance := ParseRotation(line)
		dial, zeroTouches = RotateDial(dial, direction, distance)
		if dial == 0 {
			password += 1
		}

		if includeZeroTouches {
			password += zeroTouches
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

	distance, _ = strconv.Atoi(line[1:])
	return
}

func RotateDial(dial int, direction Direction, distance int) (int, int) {
	var zeroTouches = distance / NUM_TICKS
	distance = distance % NUM_TICKS

	if direction == Right {
		dial = dial + distance
		if dial > DIAL_MAX {
			zeroTouches += 1
			dial = dial - NUM_TICKS

			if dial == 0 {
				zeroTouches -= 1
			}
		}
	} else {
		var dialStartedAtZero = dial == 0

		dial -= distance
		if dial < DIAL_MIN {
			if !dialStartedAtZero {
				zeroTouches += 1
			}
			dial = NUM_TICKS + dial
		}
	}

	return dial, zeroTouches
}
