package day07

import (
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var input []byte

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	inputCopy := make([]byte, len(input))
	_ = copy(inputCopy, input)

	inputCopy = bytes.TrimRight(inputCopy, "\n")
	var manifold = bytes.Split(inputCopy, []byte{'\n'})

	answer = TraverseManifold(manifold)

	return answer, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	input = bytes.TrimRight(input, "\n")
	var manifold = bytes.Split(input, []byte{'\n'})

	_ = TraverseManifold(manifold)
	answer = TraverseQuantumManifold(manifold)

	return answer, time.Since(start)
}

const (
	ENTRANCE = 'S'
	SPLITTER = '^'
	EMPTY    = '.'
	BEAM     = '|'
)

func TraverseManifold(lines [][]byte) (numSplits int) {
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == ENTRANCE {
				// Entrance should be first line so shouldn't need to check bounds here
				lines[r+1][c] = BEAM
			} else if lines[r][c] == SPLITTER {
				if lines[r-1][c] == BEAM {
					// Split beam
					markBeam(lines, r, c-1)
					markBeam(lines, r, c+1)
					markBeam(lines, r+1, c-1)
					markBeam(lines, r+1, c+1)
					numSplits += 1
				}
			} else if lines[r][c] == BEAM {
				markBeam(lines, r+1, c)
			}
		}
	}

	return
}

func TraverseQuantumManifold(lines [][]byte) (numPaths int) {
	var paths = make([][]int, len(lines))
	for i := range paths {
		paths[i] = make([]int, len(lines[i]))
	}

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			var val = lines[r][c]
			if val == ENTRANCE {
				paths[r][c] = 1
			} else if val == BEAM {
				// Copy value from above
				paths[r][c] += paths[r-1][c]
			} else if val == SPLITTER {
				// Copy value from above
				var splitVal = getVal(paths, r-1, c)
				paths[r][c-1] += splitVal
				paths[r][c+1] += splitVal
			}
		}
	}

	//for r := 0; r < len(lines); r++ {
	//	for c := 0; c < len(lines[r]); c++ {
	//		var val = lines[r][c]
	//		if val == EMPTY {
	//			fmt.Print(".")
	//		} else if val == SPLITTER {
	//			fmt.Print("^")
	//		} else if val == BEAM {
	//			fmt.Printf("%X", paths[r][c])
	//		}
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()

	var r = len(lines) - 1
	for c := 0; c < len(lines[r]); c++ {
		numPaths += paths[r][c]
	}

	return
}

func markBeam(lines [][]byte, row, col int) {
	if row < 0 || row >= len(lines) {
		return
	}

	if col < 0 || col >= len(lines[row]) {
		return
	}

	if lines[row][col] == EMPTY {
		lines[row][col] = BEAM
	}
}

func markEmpty(lines [][]byte, row, col int) int {
	if row < 0 || row >= len(lines) {
		return 0
	}

	if col < 0 || col >= len(lines[row]) {
		return 0
	}

	if lines[row][col] == BEAM {
		lines[row][col] = EMPTY
		return 1
	}

	return 0
}

func getVal(paths [][]int, row, col int) int {
	if row < 0 || row >= len(paths) {
		return 0
	}

	if col < 0 || col >= len(paths[row]) {
		return 0
	}

	return paths[row][col]
}
