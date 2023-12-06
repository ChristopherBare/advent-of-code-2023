package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

// Map of spelled-out numbers to their digit representation
var spelledOutNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func extractRealDigits(line string) (firstDigit, lastDigit string) {
	words := strings.Fields(line)
	for _, word := range words {
		fmt.Println(word)
		// Check if the word is a spelled-out number
		if digit, ok := spelledOutNumbers[word]; ok {
			fmt.Println("digit", digit)
			fmt.Println("ok", ok)
			fmt.Println("spelledOutNumbers[word]", spelledOutNumbers[word])
			// If it's a spelled-out number, replace it with its digit representation
			line = strings.Replace(line, word, digit, -1)
		}
	}
	fmt.Println("replaced strings line", line)
	// Extract the first and last "digits" after replacement
	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		}
	}

	return firstDigit, lastDigit
}

func sumNumbersInLine(line string) int {
	firstDigit, lastDigit := extractRealDigits(line)

	// Calculate the calibration value
	if firstDigit != "" && lastDigit != "" {
		calibrationValue, _ := strconv.Atoi(firstDigit + lastDigit)
		return calibrationValue
	}
	return 0
}

func main() {
	filePath := "1/example2.txt"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	fmt.Println(lines)

	totalSum := 0

	// Process each line in the file
	for _, line := range lines {
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		if line != "" {
			lineSum := sumNumbersInLine(line)
			//fmt.Println(lineSum)
			totalSum += lineSum
		}
	}

	fmt.Printf("Total sum of all numbers in 2-digit format in the file: %d\n", totalSum)
}
