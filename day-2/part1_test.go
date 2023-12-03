package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartOne(t *testing.T) {
	expectedOutput := []bool{true, true, false, false, true}
	output := make([]bool, 0)
	gc := &GameConfig{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	inputFile, err := os.Open("testInput.txt")
	require.NoError(t, err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		gameLine := fileScanner.Text()
		gameLineTokens := strings.Split(gameLine, ":")

		possible := isGamePossible(gameLineTokens[1], gc)
		output = append(output, possible)
	}

	assert.Equal(t, expectedOutput, output)
}
