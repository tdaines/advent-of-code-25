package day08_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day08"
)

func TestParseJunctionBoxes(t *testing.T) {
	var lines = []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	var boxes = day08.ParseJunctionBoxes(lines)

	assert.Equal(t, 20, len(boxes))

	assert.Equal(t, 162, boxes[0].X)
	assert.Equal(t, 817, boxes[0].Y)
	assert.Equal(t, 812, boxes[0].Z)
	assert.Equal(t, "162,817,812", boxes[0].Key)
}

func TestCreatePairs(t *testing.T) {
	var lines = []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	var boxes = day08.ParseJunctionBoxes(lines)
	var pairs = day08.CreatePairs(boxes)

	assert.Equal(t, 190, len(pairs))

	assert.Equal(t, "162,817,812", pairs[0].Left.Key)
	assert.Equal(t, "425,690,689", pairs[0].Right.Key)

	assert.Equal(t, "162,817,812", pairs[1].Left.Key)
	assert.Equal(t, "431,825,988", pairs[1].Right.Key)
}

func TestCreateCircuits(t *testing.T) {
	var lines = []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	var boxes = day08.ParseJunctionBoxes(lines)
	var pairs = day08.CreatePairs(boxes)

	var circuits = day08.CreateCircuits(boxes, pairs, 10)

	assert.Equal(t, 11, len(circuits))
	assert.Equal(t, 5, circuits[0].Cardinality())
	assert.Equal(t, 4, circuits[1].Cardinality())
	assert.Equal(t, 2, circuits[2].Cardinality())
	assert.Equal(t, 2, circuits[3].Cardinality())
}

func TestCreateFullCircuit(t *testing.T) {
	var lines = []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	var boxes = day08.ParseJunctionBoxes(lines)
	var pairs = day08.CreatePairs(boxes)

	var x1, x2 = day08.CreateFullCircuit(boxes, pairs)

	assert.Equal(t, 216, x1)
	assert.Equal(t, 117, x2)
}
