package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	findValidMuls()
}

func findValidMuls() {
	var memory = ReadData()

	r := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	matches := r.FindAllStringSubmatch(memory, -1)

	var total int = 0
	for _, mul := range matches {
		var mul1, _ = strconv.Atoi(mul[1])
		var mul2, _ = strconv.Atoi(mul[2])
		total += (mul1 * mul2)
	}

	fmt.Println(total)
}
