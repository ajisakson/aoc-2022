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
				// draw
				runningScore += 1 + 3
			case "Y":
				// win
				runningScore += 2 + 6
			case "Z":
				// loss
				runningScore += 3 + 0
			}
		case "B":
			switch myPlay {
			case "X":
				// loss
				runningScore += 1 + 0
			case "Y":
				// draw
				runningScore += 2 + 3
			case "Z":
				// win
				runningScore += 3 + 6
			}
		case "C":
			switch myPlay {
			case "X":
				// win
				runningScore += 1 + 6
			case "Y":
				// loss
				runningScore += 2 + 0
			case "Z":
				// draw
				runningScore += 3 + 3
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
