package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const test = true

// returns a list of totals, and a list of lists of operands
func ReadData() ([]int, [][]int) {
	var filepath string
	if test {
		filepath = "test-callibrations.txt"
	} else {
		filepath = "callibrations.txt"
	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var calibrations []int
	var operands [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var vals = strings.Split(scanner.Text(), ":")

		// convert all operands to ints and add them
		var operandList []int
		for _, operand := range strings.Fields(vals[1]) {
			var operandInt, _ = strconv.Atoi(operand)
			operandList = append(operandList, operandInt)
		}

		var calibInt, _ = strconv.Atoi(vals[0])
		operands = append(operands, operandList)
		calibrations = append(calibrations, calibInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return calibrations, operands
}
