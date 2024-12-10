package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	X, Y int
}

func main() {
	// Part 1
	fmt.Println("result for part 1 is:", partOne())
	// Part 2
	fmt.Println("result for part 2 is:", partTwo())
}

func partOne() int {
	data := readInput("data.txt")
	grid := mapEntities(data)
	width, height := len(data[0]), len(data)

	uniqueLocations := make(map[Position]bool)
	for _, entityList := range grid {
		for _, loc := range calculateLocations(entityList, width, height) {
			uniqueLocations[loc] = true
		}
	}

	return len(uniqueLocations)
}

func partTwo() int {
	data := readInput("data.txt")
	grid := mapEntities(data)
	width, height := len(data[0]), len(data)

	locations := findLocations(grid, width, height)
	data = updateGrid(data, locations)

	return countOccupiedCells(data)
}

func mapEntities(data []string) map[rune][]Position {
	grid := make(map[rune][]Position)
	for y, row := range data {
		for x, symbol := range row {
			if symbol != '.' {
				grid[symbol] = append(grid[symbol], Position{X: x, Y: y})
			}
		}
	}
	return grid
}

func calculateLocations(entityList []Position, width, height int) []Position {
	var locations []Position
	for i, primary := range entityList {
		for j, secondary := range entityList {
			if i != j {
				inverseX, inverseY := 2*primary.X-secondary.X, 2*primary.Y-secondary.Y
				if isValid(inverseX, inverseY, width, height) {
					locations = append(locations, Position{X: inverseX, Y: inverseY})
				}
			}
		}
	}
	return locations
}

func findLocations(grid map[rune][]Position, width, height int) []Position {
	var locations []Position
	for _, entityList := range grid {
		for i := 0; i < len(entityList); i++ {
			primary := entityList[i]
			for j := 0; j < len(entityList); j++ {
				if i != j {
					secondary := entityList[j]
					deltaX, deltaY := secondary.X-primary.X, secondary.Y-primary.Y
					antiX, antiY := secondary.X+deltaX, secondary.Y+deltaY
					for isValid(antiX, antiY, width, height) {
						locations = append(locations, Position{X: antiX, Y: antiY})
						antiX += deltaX
						antiY += deltaY
					}
				}
			}
		}
	}
	return locations
}

func isValid(x, y, width, height int) bool {
	return 0 <= x && x < width && 0 <= y && y < height
}

func updateGrid(data []string, points []Position) []string {
	for _, point := range points {
		data[point.Y] = modifyCharAtIndex(data[point.Y], point.X, '#')
	}
	return data
}

func countOccupiedCells(data []string) int {
	count := 0
	for _, row := range data {
		for _, symbol := range row {
			if symbol != '.' {
				count++
			}
		}
	}
	return count
}

func modifyCharAtIndex(str string, index int, newChar rune) string {
	runes := []rune(str)
	runes[index] = newChar
	return string(runes)
}

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
