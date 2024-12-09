package main

import (
	"bufio"
	. "dec8/myTypes"
	"fmt"
	"os"
)

const test = true

// returns a list of totals, and a list of lists of operands
func ReadData() map[byte][]Coordinate {
	var antennas map[byte][]Coordinate = map[byte][]Coordinate{}
	var filepath string
	if test {
		filepath = "test-antennas.txt"
	} else {
		filepath = "antennas.txt"
	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var calibrations []int
	var operands [][]int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var row, col int
	for scanner.Scan() {
		var rune = scanner.Bytes()[0]

		if rune != '.' {
			// we have found an antenna
			antennas[rune] = append(antennas[rune], Coordinate{Row: row, Col: col})
		}

		col++
		if rune == '\n' {
			// m = append(m, mLine)

			row++
			col = 0
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return antennas
}
