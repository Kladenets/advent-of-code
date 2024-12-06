package main

import "fmt"

func main() {
	var m = ReadMap()

	prettyPrintMap(m)
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
