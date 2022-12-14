package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// could likely just use string -> char -> to ASCII value conversion here then subtract the offset to get the values to match, but
	// the map here should save on some type conversion and may prove faster and more readable
	priority := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}

	sumOfPriorities := 0

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open file")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var contentsA string
	var contentsB string
	var contentsC string
	lineCount := 1
	for fileScanner.Scan() {

		switch lineCount {
		case 1:
			contentsA = fileScanner.Text()
		case 2:
			contentsB = fileScanner.Text()
		case 3:
			contentsC = fileScanner.Text()
			for i, c := range contentsA {
				_ = i
				if strings.ContainsRune(contentsB, c) && strings.ContainsRune(contentsC, c) {
					sumOfPriorities += priority[c]
					break
				}
			}
			lineCount = 0
		}

		lineCount++
	}

	fmt.Printf("The sum of priorities is %d", sumOfPriorities)

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close file")
	}
}
