package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open file")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")

	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close file")
	}
}

// Rock: A, X
// Paper: B, Y
// Scissors: C, Z
