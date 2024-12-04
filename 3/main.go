package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	lines := readFile("data.txt")

	wholeText := strings.Join(lines, " ")

	var factorX int
	var factorY int
	var parseX bool
	var parseY bool
	var count int

	// part 1
	for i := 0; i < len(wholeText); i++ {

		if !(i+4 < len(wholeText)) || wholeText[i:i+4] != "mul(" {
			continue
		}

		j := i + 4
		factorX = 0
		factorY = 0

		parseX = true
		for parseX && j < len(wholeText) {
			if unicode.IsDigit(rune(wholeText[j])) {
				factorX = factorX*10 + int(wholeText[j]-'0')
				j++
			} else {
				if wholeText[j] == ',' {
					parseX = false
				} else if unicode.IsLetter(rune(wholeText[j])) || unicode.IsSpace(rune(wholeText[j])) {
					break
				}
				j++
			}
		}

		parseY = true
		for parseY && j < len(wholeText) {
			if unicode.IsDigit(rune(wholeText[j])) {
				factorY = factorY*10 + int(wholeText[j]-'0')
				j++
			} else {
				if wholeText[j] == ')' {
					parseY = false
				} else if unicode.IsLetter(rune(wholeText[j])) || unicode.IsSpace(rune(wholeText[j])) {
					break
				}
				j++
			}
		}

		if !parseX && !parseY {
			count += factorX * factorY
		}
	}

	fmt.Println("sum factor:", count)

	// part 2
	mulModus := true
	count = 0
	i := 0
	for i < len(wholeText) {
		if i+4 < len(wholeText) && wholeText[i:i+4] == "do()" {
			mulModus = true
			i += 4
			continue
		}

		if i+7 < len(wholeText) && wholeText[i:i+7] == "don't()" {
			mulModus = false
			i += 7
			continue
		}

		if !mulModus {
			i++
			continue
		}

		if i+4 < len(wholeText) && wholeText[i:i+4] == "mul(" {
			i += 4
			factorX = 0
			factorY = 0

			parseX := true
			for parseX && i < len(wholeText) {
				if unicode.IsDigit(rune(wholeText[i])) {
					factorX = factorX*10 + int(wholeText[i]-'0')
					i++
				} else {
					if wholeText[i] == ',' {
						parseX = false
					} else if unicode.IsLetter(rune(wholeText[i])) || unicode.IsSpace(rune(wholeText[i])) {
						break
					}
					i++
				}
			}

			parseY := true
			for parseY && i < len(wholeText) {
				if unicode.IsDigit(rune(wholeText[i])) {
					factorY = factorY*10 + int(wholeText[i]-'0')
					i++
				} else {
					if wholeText[i] == ')' {
						parseY = false
					} else if unicode.IsLetter(rune(wholeText[i])) || unicode.IsSpace(rune(wholeText[i])) {
						break
					}
					i++
				}
			}
			if !parseX && !parseY {
				count += factorX * factorY
			}
		} else {
			i++
		}
	}
	fmt.Println("sum factor just do():", count)
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
