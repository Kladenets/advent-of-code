package main

import "fmt"

func main() {
	getSafe()
}

func getSafe() {
	var allLevels = ReadData()

	var safeCount = 0

	for _, levels := range allLevels {
		var posDirection bool
		var safe = false

		for j := 0; j < len(levels)-1; j++ {
			var diff = levels[j] - levels[j+1]
			var absDiff = max(diff, -diff)

			if j == 0 {
				posDirection = diff < 0
			} else if posDirection != (diff < 0) {
				safe = false
				break
			}

			if absDiff < 1 || absDiff > 3 {
				safe = false
				break
			} else {
				safe = true
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
