package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		print(err.Error())
		return
	}

	inputReader := bufio.NewReader(file)

	sum := 0
	for line, err := inputReader.ReadString('\n'); err == nil; line, err = inputReader.ReadString('\n') {
		print(line)
		digits := extractDigits(line)

		for _, num := range digits {
			fmt.Printf("%d ", num)
		}

		num := makeNumber(digits)
		println(num)
		sum += num
	}

	fmt.Printf("Sum is: %d\n", sum)
}

func extractDigits(s string) []int {
	results := []int{}
	spelled_digits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	buffer := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			results = append(results, int(char-'0'))
			buffer = ""
		} else {
			buffer = buffer + string(char)
			substrings := generateSubstrings(buffer)
			for _, s := range substrings {
				index := indexOf(spelled_digits, s)
				if index > -1 {
					results = append(results, index+1)
					buffer = buffer[len(buffer)-1:]
					break
				}
			}
		}
	}

	return results
}

func indexOf(items []string, item string) int {
	for index, s := range items {
		if s == item {
			return index
		}
	}
	return -1
}

func makeNumber(digits []int) int {
	if len(digits) == 0 {
		return 0
	}

	return digits[0]*10 + digits[len(digits)-1]
}

func generateSubstrings(input string) []string {
	var result []string

	for i := 0; i < len(input); i++ {
		substring := input[i:]
		result = append(result, substring)
	}

	return result
}
