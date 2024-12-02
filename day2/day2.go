package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToIntSlice(s string) []int {
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

	return numbers
}

func isReportSafe(numbers []int) bool {
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

	return true
}

func partOneSolution(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		if isReportSafe(stringToIntSlice(scanner.Text())) == true {
			sum += 1
		}
	}
	return sum
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func removeOneMakesSafe(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if isReportSafe(removeIndex(numbers, i)) {
			return true
		}
	}
	return false
}

// This is disgustingly inefficient but oh well
func partTwoSolution(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		if isReportSafe(stringToIntSlice(scanner.Text())) == true || removeOneMakesSafe(stringToIntSlice(scanner.Text())) == true {
			sum += 1
		}
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

	fmt.Println(partTwoSolution(file))
	fmt.Println(partOneSolution(file))
}
