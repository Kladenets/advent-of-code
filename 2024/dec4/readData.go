package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadData() []string {
	// filepath := "crossword.txt"
	filepath := "crossword-test.txt"

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var crossword []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		crossword = append(crossword, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// fmt.Println(crossword)
	return crossword
}
