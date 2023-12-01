package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMapping = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func readData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func getDigitFromLine(line string) int {
	var first, last int
	numberOfFindings := 0
	for _, c := range line {
		if unicode.IsDigit(c) {
			if numberOfFindings > 0 {
				last = int(c - '0')
				numberOfFindings++
			} else {
				first = int(c - '0')
				numberOfFindings++
			}
		}
	}
	if numberOfFindings == 1 {
		return first*10 + first
	}
	return first*10 + last
}

func replaceDigits(line string) string {
	actualLine := line
	smallestIndex := len(line) + 1
	var firstStringDigit string
	heighestIndex := -1
	var lastStringDigit string
	for k := range digitMapping {
		firstIndex := strings.Index(line, k)
		lastIndex := strings.LastIndex(line, k)
		if firstIndex != -1 && firstIndex < smallestIndex {
			smallestIndex = firstIndex
			firstStringDigit = k
		}
		if lastIndex != -1 && lastIndex > heighestIndex {
			heighestIndex = lastIndex
			lastStringDigit = k
		}
	}

	if smallestIndex < len(line)+1 {
		actualLine = strings.ReplaceAll(actualLine, firstStringDigit, strconv.Itoa(digitMapping[firstStringDigit]))
	}

	if heighestIndex > -1 {
		actualLine = strings.ReplaceAll(actualLine, lastStringDigit, strconv.Itoa(digitMapping[lastStringDigit]))
	}
	return actualLine
}

func Part1(filename string) int {
	sum := 0
	for _, line := range readData(filename) {
		sum += getDigitFromLine(line)
	}
	return sum
}

func Part2(filename string) int {
	sum := 0
	for _, line := range readData(filename) {
		actualLine := replaceDigits(line)

		sum += getDigitFromLine(actualLine)
	}
	return sum
}
