package day01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day01"
)

func TestRotate(t *testing.T) {
	var dial = 50

	dial, zeroTouches := day01.RotateDial(dial, day01.Left, 68)
	assert.Equal(t, 82, dial)
	assert.Equal(t, 1, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 30)
	assert.Equal(t, 52, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Right, 48)
	assert.Equal(t, 0, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 5)
	assert.Equal(t, 95, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Right, 60)
	assert.Equal(t, 55, dial)
	assert.Equal(t, 1, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 55)
	assert.Equal(t, 0, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 1)
	assert.Equal(t, 99, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 99)
	assert.Equal(t, 0, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Right, 14)
	assert.Equal(t, 14, dial)
	assert.Equal(t, 0, zeroTouches)

	dial, zeroTouches = day01.RotateDial(dial, day01.Left, 82)
	assert.Equal(t, 32, dial)
	assert.Equal(t, 1, zeroTouches)
}

func TestParseRotation(t *testing.T) {
	direction, distance := day01.ParseRotation("R5")
	assert.Equal(t, day01.Right, direction)
	assert.Equal(t, 5, distance)

	direction, distance = day01.ParseRotation("R49")
	assert.Equal(t, day01.Right, direction)
	assert.Equal(t, 49, distance)

	direction, distance = day01.ParseRotation("L42")
	assert.Equal(t, day01.Left, direction)
	assert.Equal(t, 42, distance)

	direction, distance = day01.ParseRotation("L469")
	assert.Equal(t, day01.Left, direction)
	assert.Equal(t, 469, distance)
}

func TestGetPasswordWithouZeroTouches(t *testing.T) {
	var lines = []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	var password = day01.GetPassword(lines, false)
	assert.Equal(t, 3, password)
}

func TestGetPasswordWithZeroTouches(t *testing.T) {
	var lines = []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	var password = day01.GetPassword(lines, true)
	assert.Equal(t, 6, password)
}
