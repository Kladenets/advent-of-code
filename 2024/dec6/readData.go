package main

import (
	"bufio"
	"fmt"
	"os"
)

const test = false

func ReadMap() ([][]byte, int, int) {
	var filepath string
	if test {
		filepath = "test-map.txt"
	} else {
		filepath = "map.txt"
	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var m [][]byte
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var mLine []byte
	var startRow, startColumn int
	var found bool = false
	for scanner.Scan() {
		var rune = scanner.Bytes()[0]
		if rune == '^' {
			found = true
		}

		if !found {
			startColumn++
		}

		if rune != '\n' {
			mLine = append(mLine, rune)
		} else {
			m = append(m, mLine)
			if !found {
				startRow++
				startColumn = 0
			}
			mLine = nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		m = append(m, mLine)
	}

	return m, startRow, startColumn
}
