package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type set struct {
	red, green, blue int
}

func (s set) possible() bool {
	return s.red <= 12 && s.green <= 13 && s.blue <= 14
}

type game struct {
	id   int
	sets []set
}

func (g game) possible() bool {
	for _, s := range g.sets {
		if !s.possible() {
			return false
		}
	}
	return true
}

func (g game) power() int {
	red, green, blue := -1, -1, -1
	for _, s := range g.sets {
		red = int(math.Max(float64(red), float64(s.red)))
		green = int(math.Max(float64(green), float64(s.green)))
		blue = int(math.Max(float64(blue), float64(s.blue)))
	}
	return red * green * blue
}

func readData(filename string) []game {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	games := make([]game, 0)

	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		stringSets := strings.Split(line[strings.Index(line, ":")+1:], ";")
		sets := make([]set, 0, len(stringSets))
		for _, s := range stringSets {
			newSet := set{-1, -1, -1}
			cubes := strings.Split(s, ",")
			for _, c := range cubes {
				c = strings.TrimSpace(c)
				tmp := strings.Split(c, " ")
				n, _ := strconv.Atoi(tmp[0])
				switch tmp[1] {
				default:
					panic("unknown color: " + tmp[1])
				case "red":
					newSet.red = n
				case "green":
					newSet.green = n
				case "blue":
					newSet.blue = n
				}
			}
			sets = append(sets, newSet)
		}

		games = append(games, game{id, sets})

		id++
	}
	return games
}

func Part1(filename string) int {
	sum := 0
	for _, g := range readData(filename) {
		if g.possible() {
			sum += g.id
		}
	}
	return sum
}

func Part2(filename string) int {
	power := 0
	for _, g := range readData(filename) {
		power += g.power()
	}
	return power
}
