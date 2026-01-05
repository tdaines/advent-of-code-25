package day08

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

	var boxes = ParseJunctionBoxes(lines)
	var pairs = CreatePairs(boxes)
	var circuits = CreateCircuits(boxes, pairs, 1000)

	answer = ProductOf3Largest(circuits)

	return answer, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")

	var boxes = ParseJunctionBoxes(lines)
	var pairs = CreatePairs(boxes)
	var x1, x2 = CreateFullCircuit(boxes, pairs)

	answer = x1 * x2

	return answer, time.Since(start)
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
