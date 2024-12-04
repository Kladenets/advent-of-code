package main

import (
	"fmt"
	"os"
)

func ReadData() string {
	filepath := "memory.txt"
	b, err := os.ReadFile(filepath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	return str
}
