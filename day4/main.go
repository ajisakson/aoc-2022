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

		if startA == startB || endA == startB || startA == endB || endA == endB {
			totalOverlaps++
			fmt.Println("hit1")
		} else if startB > startA && startB < endA {
			totalOverlaps++
			fmt.Println("hit2")
		} else if endB > startA && endB < endA {
			totalOverlaps++
			fmt.Println("hit3")
		} else if startA > startB && startA < endB {
			totalOverlaps++
			fmt.Println("hit4")
		}

		if err1 != nil {
			fmt.Println(err1.Error())
		}
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		if err3 != nil {
			fmt.Println(err3.Error())
		}
		if err4 != nil {
			fmt.Println(err4.Error())
		}
	}

	fmt.Printf("Total overlaps: %d", totalOverlaps)

	if err = file.Close(); err != nil {
		fmt.Print("Could not close file")
	}
}
