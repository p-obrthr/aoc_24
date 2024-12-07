package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DataEntry struct {
	firstNumber int
	restNumbers []int
}

func main() {
	// part 1
	fmt.Println("Total calibration result:", totalCalibrationResult())
	// part 2
	fmt.Println("Total calibration with ||:", totalCalibrationWithConcatenation())
}

func processLines(lines []string) []DataEntry {
	var dataEntries []DataEntry
	for _, line := range lines {
		entry := parseLine(line)
		dataEntries = append(dataEntries, entry)
	}
	return dataEntries
}

func parseLine(line string) DataEntry {
	parts := strings.Split(line, ":")
	firstNumber, _ := strconv.Atoi(parts[0])
	restNumbers := parseNumbers(parts[1])
	return DataEntry{
		firstNumber: firstNumber,
		restNumbers: restNumbers,
	}
}

func parseNumbers(s string) []int {
	var numbers []int
	parts := strings.Fields(s)
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		numbers = append(numbers, number)
	}
	return numbers
}

func totalCalibrationResult() int {
	lines := readFile("data.txt")
	dataEntries := processLines(lines)
	totalSum := 0

	for _, entry := range dataEntries {
		if isCalibrated(entry) {
			totalSum += entry.firstNumber
		}
	}
	return totalSum
}

func totalCalibrationWithConcatenation() int {
	lines := readFile("data.txt")
	dataEntries := processLines(lines)
	totalSum := 0

	for _, entry := range dataEntries {
		if isCalibratedWithConcatenation(entry) {
			totalSum += entry.firstNumber
		}
	}
	return totalSum
}

func isCalibrated(data DataEntry) bool {
	results := calculatePossibleResults(data.restNumbers)
	for _, result := range results {
		if result == data.firstNumber {
			return true
		}
	}
	return false
}

func isCalibratedWithConcatenation(data DataEntry) bool {
	results := calculatePossibleResultsWithConcatenation(data.restNumbers)
	for _, result := range results {
		if result == data.firstNumber {
			return true
		}
	}
	return false
}

func calculatePossibleResults(nums []int) []int {
	return calcRecursive(nums, 0, nums[0])
}

func calcRecursive(nums []int, index int, currentResult int) []int {
	if index == len(nums)-1 {
		return []int{currentResult}
	}

	nextNum := nums[index+1]

	results := []int{}

	results = append(results, calcRecursive(nums, index+1, currentResult+nextNum)...)
	results = append(results, calcRecursive(nums, index+1, currentResult*nextNum)...)

	return results
}

func calculatePossibleResultsWithConcatenation(nums []int) []int {
	return calcRecursiveWithConcatenation(nums, 0, nums[0])
}

func calcRecursiveWithConcatenation(nums []int, index int, currentResult int) []int {
	if index == len(nums)-1 {
		return []int{currentResult}
	}

	nextNum := nums[index+1]

	results := []int{}

	results = append(results, calcRecursiveWithConcatenation(nums, index+1, currentResult+nextNum)...)
	results = append(results, calcRecursiveWithConcatenation(nums, index+1, currentResult*nextNum)...)

	concatResult := concatNumbers(currentResult, nextNum)
	results = append(results, calcRecursiveWithConcatenation(nums, index+1, concatResult)...)

	return results
}

func concatNumbers(a, b int) int {
	concatenated := strconv.Itoa(a) + strconv.Itoa(b)
	result, _ := strconv.Atoi(concatenated)
	return result
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
