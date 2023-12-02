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
		digits := extractDigits(line)
		num := makeNumber(digits)
		sum += num
	}

	fmt.Printf("Sum is: %d\n", sum)
}

func extractDigits(s string) []int {
	results := []int{}
	for _, char := range s {
		if char >= '0' && char <= '9' {
			results = append(results, int(char-'0'))
		}
	}

	return results
}

func makeNumber(digits []int) int {
	if len(digits) == 0 {
		return 0
	}

	return digits[0]*10 + digits[len(digits)-1]
}
