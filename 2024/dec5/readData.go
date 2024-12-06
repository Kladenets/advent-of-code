package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPageRules() [][]bool {
	filepath := "page-rules.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	var xs, ys []int
	var maxX, maxY int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var vals = strings.Split(scanner.Text(), "|")
		var ruleX, _ = strconv.Atoi(vals[0])
		if ruleX > maxX {
			maxX = ruleX
		}

		var ruleY, _ = strconv.Atoi(vals[1])
		if ruleY > maxY {
			maxY = ruleY
		}

		xs = append(xs, ruleX)
		ys = append(ys, ruleY)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var rules = make([][]bool, maxY+1)
	for i := 0; i < len(rules); i++ {
		rules[i] = make([]bool, maxX+1)
	}

	for i := 0; i < len(xs); i++ {
		rules[xs[i]][ys[i]] = true
	}

	return rules
}

func ReadUpdates() [][]int {
	filepath := "updates.txt"
	// filepath := "test-updates.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var updates [][]int

	for scanner.Scan() {
		var vals = strings.Split(scanner.Text(), ",")
		var update []int

		for _, val := range vals {
			var valInt, _ = strconv.Atoi(val)
			update = append(update, valInt)
		}

		updates = append(updates, update)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return updates
}
