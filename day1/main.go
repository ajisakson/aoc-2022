package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var mostCalories int = 0
var secondMost int = 0
var thirdMost int = 0
var currentCalories int = 0
var lineValue int = 0

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open file")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if currentCalories > mostCalories {
				thirdMost = secondMost
				secondMost = mostCalories
				mostCalories = currentCalories
			} else if currentCalories > secondMost {
				thirdMost = secondMost
				secondMost = currentCalories
			} else if currentCalories > thirdMost {
				thirdMost = currentCalories
			}
			currentCalories = 0
		} else {
			lineValue, err = strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Printf("err found!")
			}
			currentCalories += lineValue
		}
	}

	if errors.Is(err, io.EOF) {
		if currentCalories > mostCalories {
			mostCalories = currentCalories
			currentCalories = 0
		}
	}

	topThreeTotal := mostCalories + secondMost + thirdMost
	fmt.Printf("%d is the most calories", topThreeTotal)

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close file")
	}
}
