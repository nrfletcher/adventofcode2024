package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(s string) bool {
	numbersAsStrings := strings.Fields(s)
	var numbers []int
	for i := 0; i < len(numbersAsStrings); i++ {
		val, err := strconv.Atoi(numbersAsStrings[i])
		if err != nil {
			fmt.Println("Error converting")
			os.Exit(1)
		}
		numbers = append(numbers, val)
	}

	isIncreasing := false
	if numbers[0] < numbers[1] {
		isIncreasing = true
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] && isIncreasing {
			return false
		}
		if numbers[i] < numbers[i+1] && !isIncreasing {
			return false
		}
		if math.Abs(float64(numbers[i])-float64(numbers[i+1])) > 3 || math.Abs(float64(numbers[i])-float64(numbers[i+1])) == 0 {
			return false
		}
	}

	fmt.Printf("Valid %s \n", s)
	return true
}

func partOneSolution(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		if isReportSafe(scanner.Text()) == true {
			sum += 1
		}
	}
	return sum
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
