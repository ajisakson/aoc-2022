package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	totalOverlaps := 0

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open file")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ranges := strings.Split(fileScanner.Text(), ",")
		elfARange := strings.Split(ranges[0], "-")
		elfBRange := strings.Split(ranges[1], "-")
		startA, err1 := strconv.Atoi(elfARange[0])
		endA, err2 := strconv.Atoi(elfARange[1])
		startB, err3 := strconv.Atoi(elfBRange[0])
		endB, err4 := strconv.Atoi(elfBRange[1])

		if startA <= startB && endA >= endB {
			totalOverlaps++
		} else if startB <= startA && endB >= endA {
			totalOverlaps++
		}

		if err1 != nil {
			fmt.Print(err1.Error())
		}
		if err2 != nil {
			fmt.Print(err2.Error())
		}
		if err3 != nil {
			fmt.Print(err3.Error())
		}
		if err4 != nil {
			fmt.Print(err4.Error())
		}
	}

	fmt.Printf("Total containments: %d", totalOverlaps)

	if err = file.Close(); err != nil {
		fmt.Print("Could not close file")
	}
}
