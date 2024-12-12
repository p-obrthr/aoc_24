package main

import (
	"bufio"
	"fmt"
	"os"
)

type vec [2]int

func main() {
	// part 1
	fmt.Println("distinct X:", findX())
	// part 2
	fmt.Println("loop:", findLoops())
}

func findX() int {
	lines := readFile("data.txt")
	var filteredLines []string
	for _, line := range lines {
		if len(line) > 0 {
			filteredLines = append(filteredLines, line)
		}
	}

	width, height := 0, len(filteredLines)
	for _, line := range filteredLines {
		if len(line) > width {
			width = len(line)
		}
	}

	grid := make([][]rune, height)
	for y, line := range filteredLines {
		grid[y] = make([]rune, width)
		for x := 0; x < len(line); x++ {
			grid[y][x] = rune(line[x])
		}
	}

	var start vec
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				start = vec{x, y}
				grid[y][x] = '.'
				break
			}
		}
	}

	visited := make(map[vec]bool)
	dir := vec{0, -1}
	current := start

	for {
		visited[current] = true
		next := vec{current[0] + dir[0], current[1] + dir[1]}
		if next[0] < 0 || next[1] < 0 || next[0] >= width || next[1] >= height {
			break
		}
		for next[0] >= 0 && next[1] >= 0 && next[0] < width && next[1] < height && grid[next[1]][next[0]] == '#' {
			dir = vec{-dir[1], dir[0]}
			next = vec{current[0] + dir[0], current[1] + dir[1]}
		}
		current = next
	}
	return len(visited)
}

func findLoops() int {
	lines := readFile("data.txt")
	var filteredLines []string
	for _, line := range lines {
		if len(line) > 0 {
			filteredLines = append(filteredLines, line)
		}
	}

	width, height := 0, len(filteredLines)
	for _, line := range filteredLines {
		if len(line) > width {
			width = len(line)
		}
	}

	grid := make([][]rune, height)
	for y, line := range filteredLines {
		grid[y] = make([]rune, width)
		for x := 0; x < len(line); x++ {
			grid[y][x] = rune(line[x])
		}
	}

	var start vec
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				start = vec{x, y}
				grid[y][x] = '.'
				break
			}
		}
	}

	visited := make(map[vec]bool)
	dir := vec{0, -1}
	current := start

	for {
		visited[current] = true
		next := vec{current[0] + dir[0], current[1] + dir[1]}
		if next[0] < 0 || next[1] < 0 || next[0] >= width || next[1] >= height {
			break
		}
		for next[0] >= 0 && next[1] >= 0 && next[0] < width && next[1] < height && grid[next[1]][next[0]] == '#' {
			dir = vec{-dir[1], dir[0]}
			next = vec{current[0] + dir[0], current[1] + dir[1]}
		}
		current = next
	}

	loops := 0
	for v := range visited {
		if v == start {
			continue
		}
		grid[v[1]][v[0]] = '#'
		if checkLoop(grid, width, height, start) {
			loops++
		}
		grid[v[1]][v[0]] = '.'
	}
	return loops
}

func checkLoop(grid [][]rune, width, height int, start vec) bool {
	visited := make(map[vec]vec)
	dir := vec{0, -1}
	current := start

	for {
		if visited[current] == dir {
			return true
		}
		visited[current] = dir
		next := vec{current[0] + dir[0], current[1] + dir[1]}
		if next[0] < 0 || next[1] < 0 || next[0] >= width || next[1] >= height {
			break
		}
		for next[0] >= 0 && next[1] >= 0 && next[0] < width && next[1] < height && grid[next[1]][next[0]] == '#' {
			dir = vec{-dir[1], dir[0]}
			next = vec{current[0] + dir[0], current[1] + dir[1]}
		}
		current = next
	}
	return false
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}
