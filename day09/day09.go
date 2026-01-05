package day09

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

//go:embed input.txt
var input string

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")
	answer = FindLargestRect(lines)

	return answer, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	//var lines = strings.Split(input, "\n")

	answer = 0

	return answer, time.Since(start)
}

func FindLargestRect(lines []string) int {
	var area = 0
	var coords = ParseCoords(lines)

	for i := 0; i < len(coords); i++ {
		var coord1 = coords[i]
		for j := i + 1; j < len(lines); j++ {
			var coord2 = coords[j]
			area = max(area, CalcArea(coord1.Row, coord1.Col, coord2.Row, coord2.Col))
		}
	}

	return area
}

type Coord struct {
	Row, Col int
}

func ParseCoords(lines []string) []Coord {
	var coords = make([]Coord, 0, len(lines))

	for _, line := range lines {
		coords = append(coords, ParseCoord(line))
	}

	return coords
}

func ParseCoord(line string) Coord {
	var parts = strings.Split(line, ",")

	var col, _ = strconv.Atoi(parts[0])
	var row, _ = strconv.Atoi(parts[1])

	return Coord{row, col}
}

func CalcArea(r1, c1, r2, c2 int) int {
	var height = c1 - c2
	if height < 0 {
		height *= -1
	}

	height += 1

	var width = r1 - r2
	if width < 0 {
		width *= -1
	}

	width += 1

	return height * width
}

type JunctionBox struct {
	Key string
	X   int
	Y   int
	Z   int
}

type JunctionBoxPair struct {
	Left     JunctionBox
	Right    JunctionBox
	Distance float64
}

func CreatePairs(boxes []JunctionBox) (pairs []JunctionBoxPair) {
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, createPair(boxes[i], boxes[j]))
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance < pairs[j].Distance
	})

	return
}

func createPair(left, right JunctionBox) JunctionBoxPair {
	var xDist = left.X - right.X
	var yDist = left.Y - right.Y
	var zDist = left.Z - right.Z
	var distance = math.Sqrt(float64(xDist*xDist) + float64(yDist*yDist) + float64(zDist*zDist))

	return JunctionBoxPair{
		Left:     left,
		Right:    right,
		Distance: distance,
	}
}

func ParseJunctionBoxes(lines []string) (boxes []JunctionBox) {
	for _, line := range lines {
		if len(line) > 0 {
			boxes = append(boxes, parseJunctionBox(line))
		}
	}

	return
}

func parseJunctionBox(line string) JunctionBox {
	// 162,817,812
	var parts = strings.Split(line, ",")
	var x, _ = strconv.Atoi(parts[0])
	var y, _ = strconv.Atoi(parts[1])
	var z, _ = strconv.Atoi(parts[2])

	return JunctionBox{
		Key: line,
		X:   x,
		Y:   y,
		Z:   z,
	}
}

func CreateCircuits(boxes []JunctionBox, pairs []JunctionBoxPair, numCircuits int) []mapset.Set[string] {
	var circuits = mapset.NewSet[mapset.Set[string]]()

	// Make a circuit for each junction box
	for _, box := range boxes {
		var circuit = mapset.NewSet[string]()
		circuit.Add(box.Key)
		circuits.Add(circuit)
	}

	// Create or combine circuits for each pair
	for i := range numCircuits {
		var pair = pairs[i]
		var left = findCircuit(pair.Left.Key, circuits)
		circuits.Remove(left)

		var right = findCircuit(pair.Right.Key, circuits)
		circuits.Remove(right)

		var combined = left.Union(right)
		circuits.Add(combined)
	}

	// Put all the circuits into a list so we can sort them by number of junction boxes
	var circuitList = make([]mapset.Set[string], 0, circuits.Cardinality())
	for circuit := range mapset.Elements(circuits) {
		circuitList = append(circuitList, circuit)
	}

	sort.Slice(circuitList, func(i, j int) bool {
		return circuitList[i].Cardinality() > circuitList[j].Cardinality()
	})

	return circuitList
}

func findCircuit(key string, circuits mapset.Set[mapset.Set[string]]) mapset.Set[string] {
	for circuit := range mapset.Elements(circuits) {
		if circuit.Contains(key) {
			return circuit
		}
	}

	return mapset.NewSet[string]()
}

func ProductOf3Largest(circuits []mapset.Set[string]) int {
	var product = 1

	for i := range 3 {
		product *= circuits[i].Cardinality()
	}

	return product
}

func CreateFullCircuit(boxes []JunctionBox, pairs []JunctionBoxPair) (int, int) {
	var circuits = mapset.NewSet[mapset.Set[string]]()

	// Make a circuit for each junction box
	for _, box := range boxes {
		var circuit = mapset.NewSet[string]()
		circuit.Add(box.Key)
		circuits.Add(circuit)
	}

	var leftJunction JunctionBox
	var rightJunction JunctionBox

	for i := 0; i < len(pairs); i++ {
		leftJunction = pairs[i].Left
		rightJunction = pairs[i].Right

		var leftCircuit = findCircuit(leftJunction.Key, circuits)
		circuits.Remove(leftCircuit)

		var rightCircuit = findCircuit(rightJunction.Key, circuits)
		circuits.Remove(rightCircuit)

		var combined = leftCircuit.Union(rightCircuit)
		circuits.Add(combined)

		if circuits.Cardinality() == 1 {
			break
		}
	}

	return leftJunction.X, rightJunction.X
}
