package day4

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type card struct {
	id                        int
	winningNumbers, myNumbers []int
}

func (c card) hits() int {
	hits := 0
	for _, myNumber := range c.myNumbers {
		for _, number := range c.winningNumbers {
			if myNumber == number {
				hits++
			}
		}
	}
	return hits
}

func (c card) points() int {
	return int(math.Pow(2.0, float64(c.hits()-1)))
}

func convertNumbers(str []string) []int {
	numbers := make([]int, 0, len(str))
	for _, s := range str {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}

func readData(filename string) []card {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	cards := make([]card, 0)

	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		allNumbers := strings.Split(line[strings.Index(line, ":")+1:], "|")
		winningNumbersStr := strings.Split(strings.TrimSpace(allNumbers[0]), " ")
		myNumbersStr := strings.Split(strings.TrimSpace(allNumbers[1]), " ")

		winningNumbers := convertNumbers(winningNumbersStr)
		myNumbers := convertNumbers(myNumbersStr)

		cards = append(cards, card{id, winningNumbers, myNumbers})
		id++
	}
	return cards
}

func Part1(filename string) int {
	sum := 0
	for _, c := range readData(filename) {
		sum += c.points()
	}
	return sum
}

func Part2(filename string) int {
	cardCount := make(map[int]int)

	cards := readData(filename)
	for _, c := range cards {
		cardCount[c.id]++
		for i := 1; i <= c.hits(); i++ {
			if c.id+i <= len(cards) {
				cardCount[c.id+i] += cardCount[c.id]
			}
		}
	}

	sum := 0
	for _, count := range cardCount {
		sum += count
	}
	return sum
}
