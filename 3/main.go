package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isValidCoordinate(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func sumPartNumbers(schematic []string) int {
	symbols := map[byte]bool{'*': true, '#': true, '+': true, '$': true}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	rows := len(schematic)
	cols := len(schematic[0])
	partSum := 0

	for rowIdx, row := range schematic {
		for colIdx, char := range row {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				for _, dir := range directions {
					newRow := rowIdx + dir[0]
					newCol := colIdx + dir[1]
					if isValidCoordinate(newRow, newCol, rows, cols) && symbols[schematic[newRow][newCol]] {
						partSum += digit
						break
					}
				}
			}
		}
	}

	return partSum
}

func readSchematicFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Remove leading/trailing whitespaces
	for i, line := range lines {
		lines[i] = line[1 : len(line)-1]
	}

	return lines, nil
}

func main() {
	// Path to your engine schematic file
	filePath := "3/test.txt"

	// Read the engine schematic from file
	schematicLines, err := readSchematicFromFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return
	}

	// Calculate the sum of part numbers
	totalSum := sumPartNumbers(schematicLines)
	fmt.Println("Sum of all part numbers in the engine schematic:", totalSum)
}
