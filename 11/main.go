package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part 1
	fmt.Println("result part 1 is:", partOne())
	// part 2
	fmt.Println("result part 2 is:", partTwo())
}

func partOne() int {
	return calc("data.txt", 25)
}

func partTwo() int {
	return calc("data.txt", 75)
}

func calc(filename string, iterations int) int {
	data := readFile(filename)
	list := getLists(data)[0]

	stones := initStones(list)

	for i := 0; i < iterations; i++ {
		stones = blink(stones)
	}

	numberOfStones := 0
	for _, count := range stones {
		numberOfStones += count
	}

	return numberOfStones
}

func initStones(list []int) map[int]int {
	stones := make(map[int]int)
	for _, stone := range list {
		stones[stone] = 1
	}
	return stones
}

func blink(stones map[int]int) map[int]int {
	tempStones := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			tempStones[1] += count
			continue
		}

		stoneStr := strconv.Itoa(stone)
		stoneLen := len(stoneStr)

		if stoneLen%2 == 0 {
			left, _ := strconv.Atoi(stoneStr[:stoneLen/2])
			right, _ := strconv.Atoi(stoneStr[stoneLen/2:])
			tempStones[left] += count
			tempStones[right] += count
		} else {
			tempStones[stone*2024] += count
		}
	}

	return tempStones
}

func getLists(lines []string) [][]int {
	var arrays [][]int

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == '\t' || r == ','
		})

		var numbers []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}
		arrays = append(arrays, numbers)
	}
	return arrays
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
