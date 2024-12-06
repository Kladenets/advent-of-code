package main

import (
	"fmt"
)

func main() {
	getValidUpdateMiddleNumberSum()
}

func getValidUpdateMiddleNumberSum() {
	var rules = ReadPageRules()
	var updates = ReadUpdates()
	var middleSum int = 0

	for _, update := range updates {
		var goodUpdate = true
		for x := 0; x < len(update); x++ {
			for y := 0; y < len(update); y++ {
				if x == y {
					continue
				}

				if rules[update[x]][update[y]] {
					if x > y {
						var temp = update[x]
						update[x] = update[y]
						update[y] = temp
						goodUpdate = false
						x = 0
						continue
					}
				}
			}
		}

		if !goodUpdate {
			middleSum += update[(len(update) / 2)]
		}
	}

	fmt.Println(middleSum)
}
