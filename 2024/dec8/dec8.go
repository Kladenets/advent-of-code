package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var calibrations, allOperands = ReadData()
	// fmt.Println(calibrations, allOperands)

	var totalCalibrations int
	for i, calibration := range calibrations {
		var operands = allOperands[i]
		var allOperatorsToTest = recurse(len(operands)-1, operators)
		// prettyPrintMap(allOperatorsToTest)

		// ops is a set of operators
		for _, ops := range allOperatorsToTest {
			// loop through all of operands, testing against this set of operators
			var testTotal = operands[0]
			for j := 0; j < len(operands)-1; j++ {
				testTotal = operate(ops[j], testTotal, operands[j+1])
				// fmt.Println(testTotal, calibration)
			}

			if testTotal == calibration {
				totalCalibrations += calibration
				break
			}
		}
	}

	fmt.Println(totalCalibrations)
}

func operate(operator byte, operand1 int, operand2 int) int {
	switch {
	// +
	case operator == 43:
		return operand1 + operand2
	// *
	case operator == 42:
		return operand1 * operand2
	// |
	case operator == 124:
		var concatInt, _ = strconv.Atoi(strconv.Itoa(operand1) + strconv.Itoa(operand2))
		return concatInt
	default:
		return math.MaxInt
	}
}

var operators = []byte{'+', '*', '|'}

// operators might not need to be a param
func recurse(count int, operators []byte) [][]byte {
	if count == 1 {
		var returnOps [][]byte
		for _, op := range operators {
			var tempOpList = []byte{op}
			returnOps = append(returnOps, tempOpList)
		}
		return returnOps
	}

	// loop through the list and the result of recurse, add
	var lists [][]byte

	// i is one of the operands
	for _, op := range operators {
		// j is one of the permutations
		for _, runningOps := range recurse(count-1, operators) {
			lists = append(lists, append(runningOps, op))
		}
	}

	return lists
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
