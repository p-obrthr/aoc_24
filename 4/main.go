package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := readFile("data.txt")
	matrix := createMatrix(data)

	// part 1
	word := "XMAS"
	reverseWord := "SAMX"
	count := 0
	count += countHorizontal(matrix, word, reverseWord)
	count += countVertical(matrix, word, reverseWord)
	count += countDiagonal(matrix, word, reverseWord)
	fmt.Println("horizontal, diagonal, vertical:", count)

	// part 2
	count = 0
	word = "MAS"
	reverseWord = "SAM"
	count += countMAS(matrix, word, reverseWord)
	fmt.Println("mas: ", count)
}

func countMAS(matrix [][]rune, word, reverseWord string) int {
	count := 0

	for i := 0; i < len(matrix)-2; i++ {
		for j := 0; j < len(matrix[i])-2; j++ {
			block := [][]rune{
				{matrix[i][j], matrix[i][j+1], matrix[i][j+2]},
				{matrix[i+1][j], matrix[i+1][j+1], matrix[i+1][j+2]},
				{matrix[i+2][j], matrix[i+2][j+1], matrix[i+2][j+2]},
			}

			patternElements := []rune{
				block[0][0], block[1][1], block[2][2],
				block[2][0], block[1][1], block[0][2],
			}

			pattern1 := string([]rune{patternElements[0], patternElements[1], patternElements[2]})
			pattern2 := string([]rune{patternElements[3], patternElements[4], patternElements[5]})

			if (pattern1 == word || pattern1 == reverseWord) && (pattern2 == word || pattern2 == reverseWord) {
				count++
			}
		}
	}

	return count
}

func countHorizontal(matrix [][]rune, word, reverseWord string) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= len(matrix[i])-len(word); j++ {
			if string(matrix[i][j:j+len(word)]) == word || string(matrix[i][j:j+len(reverseWord)]) == reverseWord {
				count++
			}
		}
	}
	return count
}

func countVertical(matrix [][]rune, word, reverseWord string) int {
	count := 0
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i <= len(matrix)-len(word); i++ {
			verticalWord := ""
			verticalReverseWord := ""
			for k := 0; k < len(word); k++ {
				verticalWord += string(matrix[i+k][j])
				verticalReverseWord += string(matrix[i+k][j])
			}
			if verticalWord == word || verticalReverseWord == reverseWord {
				count++
			}
		}
	}
	return count
}

func countDiagonal(matrix [][]rune, word, reverseWord string) int {
	count := 0
	wordLength := len(word)

	for i := 0; i <= len(matrix)-wordLength; i++ {
		for j := 0; j <= len(matrix[i])-wordLength; j++ {
			diagonalWord := ""
			for z := 0; z < wordLength; z++ {
				diagonalWord += string(matrix[i+z][j+z])
			}

			if diagonalWord == word || diagonalWord == reverseWord {
				count++
			}
		}
	}

	for i := len(matrix) - 1; i >= wordLength-1; i-- {
		for j := 0; j <= len(matrix[i])-wordLength; j++ {
			diagonalWord := ""
			for z := 0; z < wordLength; z++ {
				diagonalWord += string(matrix[i-z][j+z])
			}

			if diagonalWord == word || diagonalWord == reverseWord {
				count++
			}
		}
	}

	return count
}

func createMatrix(data []string) [][]rune {
	var matrix [][]rune
	for _, line := range data {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return matrix
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
