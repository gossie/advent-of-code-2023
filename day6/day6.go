package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type game struct {
	time, recordDistance int64
}

func toNumbers(line string, kerning bool) []int64 {
	numbers := make([]int64, 0)
	if kerning {
		longNumberStr := ""
		for _, t := range strings.Split(line[strings.Index(line, ":")+1:], " ") {
			if strings.TrimSpace(t) != "" {
				longNumberStr += t
			}
		}
		time, err := strconv.ParseInt(longNumberStr, 10, 64)
		if err == nil {
			numbers = append(numbers, time)
		}
	} else {
		for _, t := range strings.Split(line[strings.Index(line, ":")+1:], " ") {
			if strings.TrimSpace(t) != "" {
				time, err := strconv.ParseInt(t, 10, 64)
				if err == nil {
					numbers = append(numbers, time)
				}
			}
		}
	}
	return numbers
}

func readData(filename string, kerning bool) []game {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	records := make([]game, 0)

	var timeLine string
	var distanceLine string

	for scanner.Scan() {
		line := scanner.Text()
		if timeLine == "" {
			timeLine = line
		} else {
			distanceLine = line
		}
	}

	times := toNumbers(timeLine, kerning)
	distances := toNumbers(distanceLine, kerning)

	if len(times) != len(distances) {
		panic("length is different")
	}

	for i := 0; i < len(times); i++ {
		records = append(records, game{time: times[i], recordDistance: distances[i]})
	}

	return records
}

func numberOfWays(r game) int {
	wins := 0
	for i := int64(0); i <= r.time; i++ {
		if (r.time-i)*i > r.recordDistance {
			wins++
		}
	}
	return wins
}

func Part1(filename string) int {
	result := 1
	for _, r := range readData(filename, false) {
		result *= numberOfWays(r)
	}
	return result
}

func Part2(filename string) int {
	result := 1
	for _, r := range readData(filename, true) {
		result *= numberOfWays(r)
	}
	return result
}
