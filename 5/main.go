package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pipedLists, commaLists := readParseDataFile("data.txt")

	// Part 1
	count := 0
	for _, commaList := range commaLists {
		if isSafe(commaList, pipedLists) {
			count += getMid(commaList)
		}
	}
	fmt.Println("count:", count)

	// Part 2
	count = 0
	for _, commaList := range commaLists {
		if !isSafe(commaList, pipedLists) {
			sortedList := orderList(pipedLists, commaList)
			mid := getMid(sortedList)
			count += mid
		}
	}
	fmt.Println("count mod:", count)
}

func isSafe(commaList []int, pipedLists [][2]int) bool {
	for i := 0; i < len(commaList)-1; i++ {
		for j := i + 1; j < len(commaList); j++ {
			x := commaList[i]
			y := commaList[j]

			for _, pair := range pipedLists {
				if (pair[0] == x && pair[1] == y) || (pair[0] == y && pair[1] == x) {
					if pair[0] == y && pair[1] == x {
						return false
					}
				}
			}
		}
	}
	return true
}

func orderList(pipedLists [][2]int, commaList []int) []int {
	sortedList := make([]int, len(commaList))
	copy(sortedList, commaList)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(sortedList)-1; i++ {
			x := sortedList[i]
			y := sortedList[i+1]

			for _, pair := range pipedLists {
				if pair[0] == y && pair[1] == x {
					sortedList[i], sortedList[i+1] = sortedList[i+1], sortedList[i]
					changed = true
				}
			}
		}
	}

	return sortedList
}

func getMid(list []int) int {
	mid := len(list) / 2
	if len(list)%2 == 0 {
		return list[mid-1]
	}
	return list[mid]
}

func readParseDataFile(filename string) ([][2]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var pipedListe [][2]int
	var commaListe [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				var pair [2]int
				fmt.Sscanf(parts[0], "%d", &pair[0])
				fmt.Sscanf(parts[1], "%d", &pair[1])
				pipedListe = append(pipedListe, pair)
			}
		} else if strings.Contains(line, ",") {
			values := strings.Split(line, ",")
			var row []int
			for _, val := range values {
				var num int
				fmt.Sscanf(val, "%d", &num)
				row = append(row, num)
			}
			commaListe = append(commaListe, row)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return pipedListe, commaListe
}
