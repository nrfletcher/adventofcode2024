package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLine() {

}

func getLineValue(s string) int {
	total := 0

	return total
}

func partOneSolution(file *os.File) int {
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sum += getLineValue(scanner.Text())
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println(partOneSolution(file))
}
