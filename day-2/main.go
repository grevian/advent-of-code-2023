package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	PartOneHandle()
	PartTwoHandle()
}

func PartOneHandle() {
	gc := &GameConfig{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	inputFile, err := os.Open("day-2/input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	possibleGameSum := 0
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		gameLine := fileScanner.Text()
		gameLineTokens := strings.Split(gameLine, ":")

		possible := isGamePossible(gameLineTokens[1], gc)
		if possible {
			var gameId int
			fmt.Sscanf(gameLineTokens[0], "Game %d", &gameId)
			possibleGameSum += gameId
		}
	}

	fmt.Println("Sum of possible winning game IDs: " + strconv.Itoa(possibleGameSum))
}

func PartTwoHandle() {
	inputFile, err := os.Open("day-2/input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	gamePowerSum := 0
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		gameLine := fileScanner.Text()
		gameLineTokens := strings.Split(gameLine, ":")

		minimumConfig := minimumCubes(gameLineTokens[1])
		gamePower := minimumConfig.Red * minimumConfig.Green * minimumConfig.Blue
		gamePowerSum += gamePower
	}

	fmt.Println("Sum of powers of all minimum configurations: " + strconv.Itoa(gamePowerSum))
}
