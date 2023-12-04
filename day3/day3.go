package day3

import (
	"bufio"
	"math"
	"os"
	"unicode"

	"github.com/gossie/aoc-utils/geometry"
)

type node struct {
	start    geometry.Point2d
	nodeType string
	number   int
	sign     string
	length   int
}

func createNumericNode(start geometry.Point2d, number int) node {
	return node{
		start,
		"num",
		number,
		"",
		int(math.Log10(float64(number))) + 1,
	}
}

func createSignNode(start geometry.Point2d, sign string) node {
	return node{
		start,
		"sign",
		0,
		sign,
		1,
	}
}

func readData(filename string) []node {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nodes := make([]node, 0)

	y := 0
	for scanner.Scan() {
		currentNumber := 0
		currentFactor := 1
		line := scanner.Text()

		for x := len(line) - 1; x >= 0; x-- {
			c := rune(line[x])
			if unicode.IsDigit(c) {
				n := int(c - '0')
				currentNumber += currentFactor * n
				currentFactor *= 10

				if x == 0 && currentNumber > 0 {
					newNode := createNumericNode(geometry.NewPoint2d(float64(x), float64(y)), currentNumber)
					nodes = append(nodes, newNode)
				}
			} else {
				if currentNumber > 0 {
					newNode := createNumericNode(geometry.NewPoint2d(float64(x+1), float64(y)), currentNumber)
					nodes = append(nodes, newNode)
				}

				if c != '.' {
					newNode := createSignNode(geometry.NewPoint2d(float64(x), float64(y)), string(c))
					nodes = append(nodes, newNode)
				}

				currentNumber = 0
				currentFactor = 1
			}
		}

		y++
	}
	return nodes
}

func Part1(filename string) int {
	sum := 0
	nodes := readData(filename)
	for _, n := range nodes {
		if n.nodeType == "num" {
			for _, s := range nodes {
				if s.nodeType == "sign" {
					if s.start.Y() >= n.start.Y()-1 && s.start.Y() <= n.start.Y()+1 && s.start.X() >= n.start.X()-1 && s.start.X() <= n.start.X()+float64(n.length) {
						sum += n.number
						break
					}
				}
			}
		}
	}
	return sum
}

func Part2(filename string) int {
	gearRatio := 0
	nodes := readData(filename)
	for _, s := range nodes {
		if s.nodeType == "sign" {
			numbers := make([]int, 0, 2)
			for _, n := range nodes {
				if n.nodeType == "num" {
					if s.start.Y() >= n.start.Y()-1 && s.start.Y() <= n.start.Y()+1 && s.start.X() >= n.start.X()-1 && s.start.X() <= n.start.X()+float64(n.length) {
						numbers = append(numbers, n.number)
					}
				}
			}
			if len(numbers) == 2 {
				gearRatio += numbers[0] * numbers[1]
			}
		}
	}
	return gearRatio
}
