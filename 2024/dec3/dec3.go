package main

import (
	"fmt"
	"sort"
)

func main() {
	getDistance()
	getSimilarity()
}

func getSimilarity() {
	var lista, listb = ReadData()

	var m map[int]int = make(map[int]int)

	for _, num := range lista {
		m[num] = 0
	}

	sim := 0
	for _, num := range listb {
		_, ok := m[num]
		if ok {
			sim += num
		}
	}

	fmt.Println("similarity: ", sim)
}

func getDistance() {
	var lista, listb = ReadData()

	sort.Ints(lista)
	sort.Ints(listb)

	distance := 0
	for i, num := range lista {
		diff := num - listb[i]

		distance += max(diff, -diff)
	}

	fmt.Println("distance: ", distance)
}
