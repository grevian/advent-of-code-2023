package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartOne() int {
	calibration := make([]int, 0)

	inputFile, err := os.Open("day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		lineScanner := bufio.NewScanner(strings.NewReader(fileScanner.Text()))
		lineScanner.Split(bufio.ScanBytes)
		first := true
		var firstSeen, lastSeen string
		for lineScanner.Scan() {
			s := lineScanner.Text()
			if strings.ContainsAny(s, "1234567890") {
				if first {
					firstSeen = s
					lastSeen = s
					first = false
				} else {
					lastSeen = s
				}
			}
		}
		var number int
		fmt.Sscanf(firstSeen+lastSeen, "%d", &number)
		calibration = append(calibration, number)
	}
	sum := 0
	for i := range calibration {
		sum += calibration[i]
	}
	return sum
}
