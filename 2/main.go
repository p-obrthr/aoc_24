package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("data.txt")
	lists := GetLists(lines)

	// part 1
	count := 0
	for _, list := range lists {
		if IsSafe(list) {
			count++
		}
	}
	fmt.Println("counter isSafe", count)

	// part 2
	count = 0
	for _, list := range lists {
		if IsSafeDumpener(list) {
			count++
		}
	}
	fmt.Println("counter isSafeDumpener", count)
}

func IsSafe(numbers []int) bool {
	direction := (numbers[1] - numbers[0]) > 0
	temp := true

	for i := 1; i < len(numbers); i++ {
		if ((numbers[i]-numbers[i-1]) > 0) != direction || !IsDistance(numbers[i], numbers[i-1]) {
			temp = false
		}
	}
	return temp
}

func IsSafeDumpener(numbers []int) bool {
	if IsSafe(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		modifiedList := append([]int{}, numbers[:i]...)
		modifiedList = append(modifiedList, numbers[i+1:]...)
		if IsSafe(modifiedList) {
			return true
		}
	}

	return false
}

func IsDistance(x, y int) bool {
	diff := math.Abs(float64(x - y))
	return diff > 0 && diff < 4
}

func GetLists(lines []string) [][]int {
	var arrays [][]int

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == '\t' || r == ','
		})

		var numbers []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("error parsing number: %v", err)
				continue
			}
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
