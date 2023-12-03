package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() int {
	calibration := make([]int, 0)

	// lol
	replacer := strings.NewReplacer(
		"oneight", "18",
		"threeight", "38",
		"fiveight", "58",
		"nineight", "98",
		"twone", "21",
		"eightwo", "82",
		"eighthree", "83",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9")

	inputFile, err := os.Open("day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Scan through the file line by line
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		// Get the current line but replace written numbers with numeric values
		currentLine := replacer.Replace(fileScanner.Text())

		// Now scan through the line byte by byte
		lineScanner := bufio.NewScanner(strings.NewReader(currentLine))
		lineScanner.Split(bufio.ScanBytes)

		// Keep track of the first and last seen numbers
		var firstSeen, lastSeen string
		first := true

		for lineScanner.Scan() {
			s := lineScanner.Text()
			if strings.ContainsAny(s, "1234567890") {
				if first {
					// Some numbers have only a single number, making it both the first and
					// last number (The calibration value still has to be a two digit number)
					lastSeen = s

					// Record the first seen number, then toggle our flag so we don't reassign it until the next line
					firstSeen = s
					first = false
				} else {
					lastSeen = s
				}
			}
		}

		// Reaching the end of the line, convert the first & last seen numbers into our calibration value
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
