package main

import (
	"bufio"
	"fmt"
	"os"
)

// Position vectors
var checkLeft = [2]int{-1, 0}
var checkRight = [2]int{}
var checkDown = [2]int{}
var checkUp = [2]int{}
var checkDiagUpLeft = [2]int{}
var checkDiagUpRight = [2]int{}
var checkDiagDownLeft = [2]int{}
var checkDiagDownRight = [2]int{}

const WORD_LENGTH = 4

func wordSearch(grid string) int {
	count := 0

	return count
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file.")
		os.Exit(1)
	}

	defer file.Close()

	input := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text()
	}

	fmt.Println(wordSearch(input))
}
