package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	var digitPositions []int
	digitMap := make(map[int]string)

	// Extract spelled-out numbers
	for word, digit := range spelledOutNumbers {
		start := 0
		for {
			pos := strings.Index(line[start:], word)
			if pos == -1 {
				break
			}
			actualPos := start + pos
			digitPositions = append(digitPositions, actualPos)
			digitMap[actualPos] = digit
			start = actualPos + len(word)
		}
	}

	// Extract numeric digits
	for pos, char := range line {
		if unicode.IsDigit(char) {
			digitMap[pos] = string(char)
			digitPositions = append(digitPositions, pos)
		}
	}

	// Sort the digitPositions
	sort.Ints(digitPositions)

	// Extract the first and last digits based on the sorted digitPositions
	var foundDigits []string
	for _, pos := range digitPositions {
		foundDigits = append(foundDigits, digitMap[pos])
	}

	if len(foundDigits) >= 1 {
		firstDigit = foundDigits[0]
		lastDigit = foundDigits[len(foundDigits)-1]
	}

	return firstDigit, lastDigit
}

func sumNumbersInLine(line string) int {
	firstDigit, lastDigit := extractRealDigits(line)
	fmt.Println("Number: ", firstDigit, lastDigit)

	// Calculate the calibration value
	if firstDigit != "" && lastDigit != "" {
		calibrationValue, _ := strconv.Atoi(firstDigit + lastDigit)
		return calibrationValue
	}
	return 0
}

func main() {
	filePath := "1/data.txt"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	totalSum := 0

	// Process each line in the file
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			lineSum := sumNumbersInLine(line)
			totalSum += lineSum
		}
	}

	fmt.Printf("Total sum of all numbers in 2-digit format in the file: %d\n", totalSum)
}
