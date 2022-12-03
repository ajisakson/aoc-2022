package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var runningScore int = 0

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
		opponentPlay := s[0]
		myPlay := s[1]

		switch opponentPlay {
		case "A":
			switch myPlay {
			case "X":
				// lose
				runningScore += 0 + 3
			case "Y":
				// draw
				runningScore += 3 + 1
			case "Z":
				// win
				runningScore += 6 + 2
			}
		case "B":
			switch myPlay {
			case "X":
				// lose
				runningScore += 0 + 1
			case "Y":
				// draw
				runningScore += 3 + 2
			case "Z":
				// win
				runningScore += 6 + 3
			}
		case "C":
			switch myPlay {
			case "X":
				// lose
				runningScore += 0 + 2
			case "Y":
				// draw
				runningScore += 3 + 3
			case "Z":
				// win
				runningScore += 6 + 1
			}
		}
	}

	fmt.Printf("Your final score is %d", runningScore)

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close file")
	}
}

// Rock: A, X
// Paper: B, Y
// Scissors: C, Z
