package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData() [][]int {
	// filepath := "levels.txt"
	filepath := "testLevels.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var allLevels [][]int

	for _, line := range lines {
		var vals = strings.Fields(line)
		var intLevels []int

		for _, val := range vals {
			var levelInt, _ = strconv.Atoi(val)
			intLevels = append(intLevels, levelInt)
		}

		allLevels = append(allLevels, intLevels)
	}

	// fmt.Println(allLevels)

	return allLevels
}
