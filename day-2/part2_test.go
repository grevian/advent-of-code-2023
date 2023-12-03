package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartTwo(t *testing.T) {
	expectedOutput := []GameConfig{
		{
			Red:   4,
			Green: 2,
			Blue:  6,
		},
		{
			Red:   1,
			Green: 3,
			Blue:  4,
		},
		{
			Red:   20,
			Green: 13,
			Blue:  6,
		},
		{
			Red:   14,
			Green: 3,
			Blue:  15,
		},
		{
			Red:   6,
			Green: 3,
			Blue:  2,
		},
	}

	inputFile, err := os.Open("testInput.txt")
	require.NoError(t, err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	index := 0
	for fileScanner.Scan() {
		gameLine := fileScanner.Text()
		gameLineTokens := strings.Split(gameLine, ":")

		minimumConfig := minimumCubes(gameLineTokens[1])
		assert.Equal(t, expectedOutput[index], minimumConfig)
		index++
	}
}
