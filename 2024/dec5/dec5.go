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
		for x, _ := range update {
			for y, _ := range update {
				if x == y {
					continue
				}

				if rules[update[x]][update[y]] {
					if x > y {
						goodUpdate = false
					}
				}
			}
		}

		if goodUpdate {
			middleSum += update[(len(update) / 2)]
		}
	}

	fmt.Println(middleSum)
}
