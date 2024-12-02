package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

/* Day 1 Part 1
 * Sort each list, and sum the differences.
 */
func partOneSolution(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var sum int
	var listOne []int
	var listTwo []int

	for scanner.Scan() {
		var s string = scanner.Text()
		var intOne bytes.Buffer
		var intTwo bytes.Buffer
		var seenSpace bool

		for i := 0; i < len(s); i++ {
			if string(s[i]) == " " {
				seenSpace = true
				continue
			}

			if seenSpace == true {
				intTwo.WriteString(string(s[i]))
			} else {
				intOne.WriteString(string(s[i]))
			}
		}

		valOne, err := strconv.Atoi(intOne.String())
		valTwo, err := strconv.Atoi(intTwo.String())

		if err != nil {
			fmt.Println("Error converting integers")
			os.Exit(1)
		}

		listOne = append(listOne, valOne)
		listTwo = append(listTwo, valTwo)
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	for i := 0; i < len(listOne); i++ {
		sum += int(math.Abs(float64(listOne[i]) - float64(listTwo[i])))
	}
	return sum
}

func similarityScore(listOne []int, listTwo []int) int {
	sum := 0

	for i := 0; i < len(listOne); i++ {
		count := 0
		value := listOne[i]
		for j := 0; j < len(listTwo); j++ {
			if listTwo[j] == value {
				count += 1
			}
		}

		sum += count * value
	}

	return sum
}

/* Day one part two.
 * Find the similarity score based off multiplication counts.
 */
func partTwoSolution(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var leftList []int
	var rightList []int

	for scanner.Scan() {
		var s string = scanner.Text()
		var intOne bytes.Buffer
		var intTwo bytes.Buffer
		var seenSpace bool

		for i := 0; i < len(s); i++ {
			if string(s[i]) == " " {
				seenSpace = true
				continue
			}

			if seenSpace == true {
				intTwo.WriteString(string(s[i]))
			} else {
				intOne.WriteString(string(s[i]))
			}
		}

		valOne, err := strconv.Atoi(intOne.String())
		valTwo, err := strconv.Atoi(intTwo.String())

		if err != nil {
			fmt.Println("Error converting integers")
			os.Exit(1)
		}

		leftList = append(leftList, valOne)
		rightList = append(rightList, valTwo)
	}

	return similarityScore(leftList, rightList)
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
