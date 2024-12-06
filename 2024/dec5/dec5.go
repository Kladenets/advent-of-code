package main

import (
	"fmt"
)

func main() {
	getValidUpdateMiddleNumberSum()
}

func getValidUpdateMiddleNumberSum() {
	var rules = ReadPageRules()
	// fmt.Println(rules[15][18])

	var updates = ReadUpdates()
	// fmt.Println(updates)

	// var goodUpdates [][]int
	var middleSum int
	for _, update := range updates {
		var goodUpdate = true
		for x, _ := range update {
			for y, _ := range update {
				if x == y {
					continue
				}

				if rules[x][y] {
					if x > y {
						goodUpdate = false
					}
				}
			}
		}

		if goodUpdate {
			// goodUpdates = append(goodUpdates, update)
			middleSum += update[(len(update) / 2)]
		}
	}

	// 5221 is too low hmm...
	fmt.Println(middleSum)
	// fmt.Println(goodUpdates)
}
