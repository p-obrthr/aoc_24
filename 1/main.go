package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("data.txt")
	left, right := splitContent(lines)

	sort.Ints(left)
	sort.Ints(right)

	// part 1
	distances := 0

	for i := 0; i < len(left); i++ {
		distances += getDistance(i, left, right)
	}

	fmt.Println("sum distances", distances)

	// part 2
	similarity := 0

	for _, element := range left {
		similarity += element * count(element, right)
	}

	fmt.Println("sum similarity", similarity)

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

func splitContent(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) > 1 {
			if leftInt, err := strconv.Atoi(parts[0]); err == nil {
				left = append(left, leftInt)
			}
			if rightInt, err := strconv.Atoi(parts[1]); err == nil {
				right = append(right, rightInt)
			}
		}
	}
	return left, right
}

func getDistance(index int, left []int, right []int) int {
	c := left[index] - right[index]
	if c < 0 {
		c = -c
	}
	return c
}

func count(number int, right []int) int {
	c := 0
	for _, n := range right {
		if n == number {
			c++
		}
	}
	return c
}
