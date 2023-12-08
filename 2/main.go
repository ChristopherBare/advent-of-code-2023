package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkPossibleGames(filename string, red, green, blue int) ([]int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var possibleGames []int
	fmt.Println(lines)
	for i, line := range lines {
		if line == "" {
			continue
		}

		// Parsing game information
		gameInfo := strings.Split(strings.Split(line, ": ")[1], "; ")
		var totalRed, totalGreen, totalBlue int

		// Calculating total cubes revealed in each game
		for _, subset := range gameInfo {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				switch {
				case strings.Contains(cube, "red"):
					totalRed += extractNumber(cube)
				case strings.Contains(cube, "green"):
					totalGreen += extractNumber(cube)
				case strings.Contains(cube, "blue"):
					totalBlue += extractNumber(cube)
				}
			}
		}
		fmt.Println("TotalRED: ", totalRed, "TotalGreen ", totalGreen, "totalBlue: ", totalBlue)
		// Checking if the total cubes match the constraints
		if totalRed <= red && totalGreen <= green && totalBlue <= blue {
			possibleGames = append(possibleGames, i+1) // Adding 1 to match game numbering
		}
	}

	return possibleGames, nil
}

func extractNumber(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return 0
	}
	return num
}

func sumGames(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	red := 12
	green := 13
	blue := 14

	possibleGames, err := checkPossibleGames("data.txt", red, green, blue)
	if err != nil {
		log.Fatal(err)
	}

	if len(possibleGames) > 0 {
		fmt.Println("Possible games:", possibleGames)
		sum := sumGames(possibleGames)
		fmt.Println("Sum: ", sum)
	} else {
		fmt.Println("No possible games found.")
	}

}
