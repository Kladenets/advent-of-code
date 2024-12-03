package main

import "fmt"

func main() {
	getSafeV2()
}

func getSafeV2() {
	// iterate through all the levels
	//   iterate through each level in a report once, and determine majority positive or negative
	//     if positive total = negative total then we can skip this report because it shouldn't be safe
	//   keep track of how many unsafe we've found, if we get to two then bail
	//   iterate through each level again, and use the same checks as before now that we now the overall direction

	// f I think I'm thinking about this wrong... I think the prompt means skip the level as if it didn't exist and
	// compare the two around it, not use it in the comparison still
	var allLevels = ReadData()

	var safeCount = 0

	for _, levels := range allLevels {
		var posDirection bool
		var posCount, negCount int
		var unSafe = 0

		for j := 0; j < len(levels)-1; j++ {
			var diff = levels[j] - levels[j+1]

			if diff < 0 {
				posCount++
			}

			if diff > 0 {
				negCount++
			}
		}

		posDirection = posCount > negCount
		// fmt.Println(posDirection)

		for j := 0; j < len(levels)-1; j++ {
			var diff = levels[j] - levels[j+1]
			var absDiff = max(diff, -diff)

			if posDirection != (diff <= 0) {
				unSafe++

			}

			if absDiff < 1 || absDiff > 3 {
				unSafe++

			}
		}

		if unSafe < 2 {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func getSafe() {
	var allLevels = ReadData()

	var safeCount = 0

	for _, levels := range allLevels {
		var posDirection bool
		var posCount, negCount int
		var safe = 0

		for j := 0; j < len(levels)-1; j++ {
			var diff = levels[j] - levels[j+1]
			var absDiff = max(diff, -diff)

			if diff <= 0 {
				posCount++
			}

			if diff >= 0 {
				negCount++
			}

			posDirection = posCount > negCount
			if posDirection != (diff < 0) {
				break
			}

			// if j == 0 {
			// 	posDirection = diff < 0
			// } else if posDirection != (diff < 0) {
			// 	break
			// }

			if absDiff < 1 || absDiff > 3 {
				break
			} else {
				safe++
			}
		}

		if safe > len(levels)-3 && (posCount <= 2 || negCount <= 2) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
