package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkPossibleGames(filename string, red, green, blue int) (int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")
	var totalPower int

	for _, line := range lines {
		if line == "" {
			continue
		}

		gameInfo := strings.Split(strings.Split(line, ": ")[1], "; ")
		var totalRed, totalGreen, totalBlue []int

		for _, subset := range gameInfo {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				switch {
				case strings.Contains(cube, "red"):
					totalRed = append(totalRed, extractNumber(cube))
				case strings.Contains(cube, "green"):
					totalGreen = append(totalGreen, extractNumber(cube))
				case strings.Contains(cube, "blue"):
					totalBlue = append(totalBlue, extractNumber(cube))
				}
			}
		}

		minRed := minCubes(totalRed)
		minGreen := minCubes(totalGreen)
		minBlue := minCubes(totalBlue)

		fmt.Println("Minred", minRed)
		fmt.Println("Mingreen", minGreen)
		fmt.Println("Minblue", minBlue)

		power := minRed * minGreen * minBlue
		totalPower += power
	}

	return totalPower, nil
}

func minCubes(totalColor []int) int {
	minNum := totalColor[0]

	// Iterate through the array to find the maximum number
	for _, num := range totalColor {
		if num > minNum {
			minNum = num
		}
	}

	return minNum
}

func extractNumber(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return 0
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	red := 12
	green := 13
	blue := 14

	totalPower, err := checkPossibleGames("data.txt", red, green, blue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of powers of minimum sets:", totalPower)

}
