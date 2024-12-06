package main

import "fmt"

func main() {
	// findXMAS()
	findXMAS3()
}

func findXMAS() {
	var crossword = ReadData()

	// i is the row number (starting at 0) of the crossword we're on
	// cwLine is the content of that row

	var noNext byte = '-'
	forwardNext := map[byte]byte{'X': 'M', 'M': 'A', 'A': 'S', 'S': noNext}
	backwardNext := map[byte]byte{'S': 'A', 'A': 'M', 'M': 'X', 'X': noNext}

	var getNextLetter = func(currentLetter byte, forward bool) byte {
		if forward {
			return forwardNext[currentLetter]
		} else {
			return backwardNext[currentLetter]
		}
	}

	var recursiveSearch func(i, j int, nextLetter byte, forward bool) int

	recursiveSearch = func(i, j int, nextLetter byte, forward bool) int {
		fmt.Println(string(nextLetter))
		if nextLetter == noNext {
			return 1
		}

		var total = 0
		for k := i; k < i+2; k++ {
			// make sure we stop if we hit the end of the crossword matrix
			if k > len(crossword)-1 {
				return 0
			}

			for l := j - 1; l < j+2; l++ {
				// if we go outside the bounds of the arrary continue
				if l < 0 {
					continue
				}

				if l > len(crossword[k])-1 {
					return 0
				}

				if k == i && l == j {
					continue
				}

				if crossword[k][l] == nextLetter {
					// if nextLetter == 'X' && k == i && l != j {
					// 	return 0
					// }

					total += recursiveSearch(k, l, getNextLetter(crossword[k][l], forward), forward)
				}
			}
		}

		return total
	}

	var total int = 0
	for i, cwLine := range crossword {
		for j := 0; j < len(cwLine); j++ {
			if cwLine[j] == 'X' {
				total += recursiveSearch(i, j, forwardNext[cwLine[j]], true)
			}
			if cwLine[j] == 'S' {
				total += recursiveSearch(i, j, backwardNext[cwLine[j]], false)
			}
		}
	}

	fmt.Println(total)
}

func findXMAS2() {
	var crossword = ReadData()

	// i is the row number (starting at 0) of the crossword we're on
	// cwLine is the content of that row

	var noNext byte = '-'
	forwardNext := map[byte]byte{'X': 'M', 'M': 'A', 'A': 'S', 'S': noNext}
	backwardNext := map[byte]byte{'S': 'A', 'A': 'M', 'M': 'X', 'X': noNext}

	var getNextLetter = func(currentLetter byte, forward bool) byte {
		if forward {
			return forwardNext[currentLetter]
		} else {
			return backwardNext[currentLetter]
		}
	}

	var recursiveSearch func(i, j int, nextLetter byte, forward bool) int

	recursiveSearch = func(i, j int, nextLetter byte, forward bool) int {
		if nextLetter == noNext {
			return 1
		}

		for k := i - 1; k < i+2; k++ {
			// make sure we don't access outside of the array
			if k < 0 {
				continue
			}

			if k > len(crossword)-1 {
				return 0
			}

			for l := j - 1; l < j+2; l++ {
				// make sure we don't go outside the row
				if l < 0 {
					continue
				}

				if l > len(crossword[k])-1 {
					return 0
				}

				if k == i && l == j {
					continue
				}

				if crossword[k][l] == nextLetter {
					return recursiveSearch(k, l, getNextLetter(crossword[k][l], forward), forward)
				}
			}
		}

		return 0
	}

	var total int = 0
	for i, cwLine := range crossword {
		for j := 0; j < len(cwLine); j++ {
			if cwLine[j] == 'X' {
				total += recursiveSearch(i, j, forwardNext[cwLine[j]], true)
			}
			// if cwLine[j] == 'S' {
			// 	total += recursiveSearch(i, j, backwardNext[cwLine[j]], false)
			// }
		}
	}

	fmt.Println(total)
}

func findXMAS3() {
	var crossword = ReadData()

	// i is the row number (starting at 0) of the crossword we're on
	// cwLine is the content of that row

	var noNext byte = '-'
	forwardNext := map[byte]byte{'X': 'M', 'M': 'A', 'A': 'S', 'S': noNext}

	var getNextLetter = func(currentLetter byte) byte {
		return forwardNext[currentLetter]
	}

	var recursiveSearch func(i, j int, nextLetter byte, xDir int, yDir int) int

	recursiveSearch = func(i, j int, nextLetter byte, xDir int, yDir int) int {
		if nextLetter == noNext {
			fmt.Println("we here")
			return 1
		}

		for k := i - 1; k < i+2; k++ {
			// make sure we don't access outside of the array
			if k < 0 {
				continue
			}

			if k > len(crossword)-1 {
				return 0
			}

			for l := j - 1; l < j+2; l++ {
				// make sure we don't go outside the row
				if l < 0 {
					continue
				}

				if l > len(crossword[k])-1 {
					return 0
				}

				if k == i && l == j {
					continue
				}

				if crossword[k][l] == nextLetter && i-k == xDir && j-l == yDir {
					fmt.Println(string(nextLetter), string(getNextLetter(crossword[k][l])))
					// fmt.Println(string(crossword[k][l]))
					return recursiveSearch(k, l, getNextLetter(crossword[k][l]), xDir, yDir)
				}
			}
		}

		return 0
	}

	var search func(i, j int, nextLetter byte) int
	var searchTotal int = 0
	search = func(i, j int, nextLetter byte) int {
		for k := i - 1; k < i+2; k++ {
			// make sure we don't access outside of the array
			if k < 0 {
				continue
			}

			if k > len(crossword)-1 {
				continue
			}

			for l := j - 1; l < j+2; l++ {
				// make sure we don't go outside the row
				if l < 0 {
					continue
				}

				if l > len(crossword[k])-1 {
					continue
				}

				// if k == i && l == j {
				// 	continue
				// }

				if crossword[k][l] == nextLetter {

					// fmt.Println(string(crossword[k][l]), i-k, j-l, string(getNextLetter(crossword[k][l])))
					// var x = i - k
					// var y = j - l
					// // fmt.Print(string(crossword[k+x][l+y]))
					// if crossword[k-x][l-y] == getNextLetter(crossword[k][l]) {
					// 	fmt.Print(string(crossword[k+x][l+y]))

					// 	if crossword[k+2*x][l+2*y] == getNextLetter(crossword[k+x][l+y]) {
					// 		fmt.Print(string(crossword[k+2*x][l+2*y]))
					// 		searchTotal += 1
					// 	}
					// }

					searchTotal += recursiveSearch(k, l, getNextLetter(crossword[k][l]), i-k, j-l)
				}
			}
		}

		return searchTotal
	}

	var total int = 0
	for i, cwLine := range crossword {
		for j := 0; j < len(cwLine); j++ {
			if cwLine[j] == 'X' {
				// for every X you find call search, which will find the direction for the search by searching all around
				// the X, then call recursive search for each match and follow the trail to see if any are valid
				total += search(i, j, forwardNext[cwLine[j]])
			}
		}
	}

	fmt.Println(total)
}
