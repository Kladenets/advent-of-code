package main

import (
	"bufio"
	"fmt"
	"os"
)

const test = true

func ReadMap() [][]byte {
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
	for scanner.Scan() {
		var rune = scanner.Bytes()[0]

		if rune != '\n' {
			mLine = append(mLine, rune)
		} else {
			m = append(m, mLine)
			mLine = nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return m
}
