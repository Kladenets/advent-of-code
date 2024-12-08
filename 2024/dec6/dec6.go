package main

import "fmt"

func main() {
	// m is the map represented in a grid, with 0,0 being the element in the left corner
	// startX, startY is the starting position of the ^
	var m, startX, startY = ReadMap()
	// fmt.Println(startX, startY)

	prettyPrintMap(m)
	var totalUniquePaths, totalNewObstacles = followGuardsPath(m, startX, startY)
	fmt.Println()
	prettyPrintMap(m)
	fmt.Println(totalUniquePaths, totalNewObstacles)
}

var turnRight = map[byte]byte{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

var turnLeft = map[byte]byte{
	'^': '<',
	'<': 'v',
	'v': '>',
	'>': '^',
}

var direction = map[byte][]int{
	'^': {-1, 0},
	'<': {0, -1},
	'v': {1, 0},
	'>': {0, 1},
}

// dir is the current guard character, one of ^, >, v, <
// 0 is straight, -1 is left, 1 is right
func turn(dir byte, turn int) byte {
	switch {
	case turn == 0:
		return dir
	case turn < 0:
		return turnLeft[dir]
	case turn > 0:
		return turnRight[dir]
	default:
		return '-'
	}
}

// part 1 we only turn right
func defaultTurn(dir byte) byte {
	return turn(dir, 1)
}

type Coordinate struct {
	row, col int
}

func followGuardsPath(m [][]byte, startRow, startColumn int) (int, int) {
	var insideMap bool = true
	// guard is the byte representing the guard, one of ^, >, v, <
	var guard = m[startRow][startColumn]
	// currentDir is an array of the current slope {x, y}
	var currentDir = direction[guard]
	// row and y track our current position in the map
	var row int = startRow
	var col int = startColumn
	// this tracks the unique path steps we took to get out of the map (crossing a path again does not count twice)
	var totalUniquePaths, totalNewObstacles int = 0, 0

	for insideMap {
		// fmt.Println(string(guard), direction, currentDir, row, col)
		// rowS = row slope, ys = col slope
		var rowS = currentDir[0]
		var colS = currentDir[1]

		//before we do anything, check if we put an obstacle in from of us and turn, will we intersect the same character
		var checkIntersection = true
		var interGuard = defaultTurn(guard)
		var interDir, firstInterGuard = direction[interGuard], guard
		var interRow, interCol = row, col
		var interPaths map[Coordinate]byte = map[Coordinate]byte{}
		var firstInterRow, firstInterCol = interRow + interDir[0], interCol + interDir[1]

		for checkIntersection {
			var interRowS = interDir[0]
			var interColS = interDir[1]
			// fmt.Println(interRow, interCol)
			var interRowTest, interColTest = interRow + interRowS, interCol + interColS
			var interCoord = Coordinate{row: interRowTest, col: interColTest}

			if (interRowTest > len(m)-1 || interRowTest < 0) || (interColTest > len(m[interRowTest])-1 || interColTest < 0) {
				checkIntersection = false
				continue
			}

			var interNextStep = m[interRowTest][interColTest]
			// fmt.Println((interRowTest == firstInterRow && interColTest == firstInterCol) && firstInterGuard == interGuard)
			// fmt.Println(interNextStep == interGuard)
			if interNextStep == interGuard || ((interRowTest == firstInterRow && interColTest == firstInterCol) && firstInterGuard == interGuard) {
				// if we find the next step that is the same as our current guard, or the starting position of this obstacle, exit the loop and count it
				totalNewObstacles++
				checkIntersection = false
				// fmt.Println(row, col)
				continue
			}

			if interNextStep == '#' || (interRowTest == firstInterRow && interColTest == firstInterCol) {
				// if we run into an obstacle we know it's not a loop, exit
				// checkIntersection = false
				interGuard = defaultTurn(interGuard)
				interDir = direction[interGuard]
				continue
			}

			var interCoordTest, _ = interPaths[interCoord]
			// if we find a loop outside of the path that we were on previously, exit the loop and count it
			if interCoordTest == interGuard {
				// we have created our own loop
				totalNewObstacles++
				checkIntersection = false
				// fmt.Println(row, col)
				continue
			}

			interPaths[interCoord] = interGuard
			interRow = interRowTest
			interCol = interColTest
		}

		// fmt.Println("here", row, col)
		// if the next step is out of bounds in either direction, we've exited the map successfully
		if (row+rowS > len(m)-1 || row+rowS < 0) || (col+colS > len(m[row+rowS])-1 || col+colS < 0) {
			insideMap = false
			m[row][col] = guard
			totalUniquePaths++
			continue
		}

		// fmt.Println("here", row, col)
		// nextStep is the step we're evaluating to see what to do
		var nextStep = m[row+rowS][col+colS]

		// if the next position is an obstacle
		if nextStep == '#' {
			// turn in place and update the map
			m[row][col] = defaultTurn(guard)

			// update the guard tracker
			guard = m[row][col]

			// update the current direction with the new direction after the turn
			currentDir = direction[guard]
			continue
		}

		// fmt.Println("here", row, col)
		// if nextStep is open and we haven't been there
		if nextStep == '.' {
			// update old position with an 'X'
			m[row][col] = guard

			// move to nextStep
			row = row + rowS
			col = col + colS

			// update the guard position
			m[row][col] = guard

			// add one to total
			totalUniquePaths++
			continue
		}

		// fmt.Println("here", row, col)
		// have we been to the nextStep already
		if nextStep == '^' || nextStep == '>' || nextStep == 'v' || nextStep == '<' {
			// move to it but do not increment totalUniquePaths, do not turn, do not update direction, do not update guard
			m[row][col] = guard
			row = row + rowS
			col = col + colS

			// update the guard position
			m[row][col] = guard

			continue
		}
	}

	return totalUniquePaths, totalNewObstacles
}

func prettyPrintMap(m [][]byte) {
	for _, row := range m {
		var prettyRow []string
		for _, byte := range row {
			prettyRow = append(prettyRow, string(byte))
		}

		fmt.Println(prettyRow)
	}
}
