package main_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStrings = []string{
	"two1nine",
	"neightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
	"tf7kndclhgjsoneoneightxcx",
}

var expectedResults = []int{
	29,
	83,
	13,
	24,
	42,
	14,
	76,
	78,
}

func TestStringConversion_Test(t *testing.T) {
	for i := range testStrings {
		currentLine := testStrings[i]
		replacer := strings.NewReplacer("oneight", "18", "one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		currentLine = replacer.Replace(currentLine)
		lineScanner := bufio.NewScanner(strings.NewReader(currentLine))
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

		assert.Equal(t, expectedResults[i], number)
	}
}
