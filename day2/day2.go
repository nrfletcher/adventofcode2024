package main

import (
	"fmt"
	"os"
)

func partOneSolution(file *os.File) int {
	return 1
}

func partTwoSolution(file *os.File) int {
	return 1
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println(partTwoSolution(file))
	fmt.Println(partOneSolution(file))
}
