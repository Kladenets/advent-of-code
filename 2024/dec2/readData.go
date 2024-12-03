package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData() ([]int, []int) {
	filepath := "lists.txt"
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

	var a, b []int

	for _, line := range lines {
		var vals = strings.Fields(line)
		var aInt, _ = strconv.Atoi(vals[0])
		var bInt, _ = strconv.Atoi(vals[1])

		a = append(a, aInt)
		b = append(b, bInt)
	}

	return a, b
}
