package main

import (
	"fmt"
	"strings"
)

type GameConfig struct {
	Red, Blue, Green int
}

func isGamePossible(gameLine string, config *GameConfig) bool {
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
		// If any of the sets of cubes pulled in this hand are greater than the game configuration, we can bail out
		if seenGreen > config.Green || seenRed > config.Red || seenBlue > config.Blue {
			return false
		}

	}
	return true
}
