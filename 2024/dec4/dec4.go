package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// findValidMuls()
	findValidMuls2()
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

func findValidMuls2() {
	var memory = ReadData()
	r := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	var goodStrs [][]string
	var splitDonts = strings.Split(memory, "don't()")
	goodStrs = append(goodStrs, r.FindAllStringSubmatch(splitDonts[0], -1)...)

	var splitDos []string
	for _, splitDont := range splitDonts[1:] {
		var tempDos = strings.Split(splitDont, "do()")

		if len(tempDos) == 1 {
			continue
		}

		splitDos = append(splitDos, tempDos[1:]...)
	}

	var strSplitDos = strings.Join(splitDos, "")
	goodStrs = append(goodStrs, r.FindAllStringSubmatch(strSplitDos, -1)...)

	var total int = 0
	for _, mul := range goodStrs {
		var mul1, _ = strconv.Atoi(mul[1])
		var mul2, _ = strconv.Atoi(mul[2])
		total += (mul1 * mul2)
	}

	fmt.Println(total)
}
