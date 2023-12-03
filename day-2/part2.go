package main

import (
	"fmt"
	"strings"
)

func minimumCubes(gameLine string) GameConfig {
	result := GameConfig{}

	subsets := strings.Split(gameLine, ";")
	for _, hand := range subsets {
		blocks := strings.Split(hand, ",")

		var seenRed, seenBlue, seenGreen int
		for _, block := range blocks {
			tokens := strings.Split(block, " ")
			numberToken := tokens[1]
			colourToken := tokens[2]
			var number int
			fmt.Sscanf(numberToken, "%d", &number)

			switch colourToken {
			case "red":
				seenRed = number
			case "blue":
				seenBlue = number
			case "green":
				seenGreen = number
			}
		}

		if seenRed > result.Red {
			result.Red = seenRed
		}
		if seenGreen > result.Green {
			result.Green = seenGreen
		}
		if seenBlue > result.Blue {
			result.Blue = seenBlue
		}

	}

	return result
}
