package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

/*
Given an index to the start of a valid "mul(" substring,
check if it ultimately is a valid mul(), and if so, return the value
*/
func parseValue(s string, idx int) int {
	ending := -1

	for i := idx; i < len(s); i++ {
		if string(s[i]) == ")" {
			ending = i
			break
		}
	}

	substr := s[idx+4 : ending]
	n_one := -1
	n_two := -1
	n_one_s := ""
	n_two_s := ""

	for i := 0; i < len(substr); i++ {
		if string(substr[i]) == "," {
			n_two_s = substr[i+1:]
			n_one_s = substr[0:i]
		}
	}

	fmt.Println(n_two_s)
	fmt.Println(n_one_s)

	n_one, err := strconv.Atoi(n_one_s)
	n_two, err_t := strconv.Atoi(n_two_s)
	if err != nil || err_t != nil {
		fmt.Println("Error converting.")
		return 0
	}

	if n_one == -1 || n_two == -1 {
		return 0
	} else {
		return n_one * n_two
	}
}

func getLineValue(s string) int {
	total := 0
	for i := 0; i < len(s)-3; i++ {
		if s[i:i+4] == "mul(" {
			total += parseValue(s, i)
		}
	}
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
