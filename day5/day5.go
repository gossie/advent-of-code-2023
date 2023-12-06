package day5

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rangeInfo struct {
	min, max int
}

type mapRange struct {
	start, destination rangeInfo
}

type data struct {
	seeds    []int
	mappings map[string][]mapRange
}

func toNumbers(numbersAsString string) []int {
	numbersStr := strings.Split(numbersAsString, " ")
	numbers := make([]int, 0, len(numbersStr))
	for _, s := range numbersStr {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}

func readData(filename string) data {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seeds := make([]int, 0)
	mappings := make(map[string][]mapRange)

	var currentMapName string
	var currentRanges []mapRange

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.HasPrefix(line, "seeds:") {
				seeds = toNumbers(line[7:])
			} else {
				if strings.Contains(line, "map:") {
					if currentMapName != "" {
						mappings[currentMapName] = currentRanges
					}
					currentMapName = strings.Split(line, " ")[0]
					currentRanges = make([]mapRange, 0)
				} else {
					mapData := toNumbers(line)
					if len(mapData) != 3 {
						panic("more than three entries in mapData")
					}

					newRange := mapRange{
						start: rangeInfo{
							min: mapData[1],
							max: mapData[1] + (mapData[2] - 1),
						},
						destination: rangeInfo{
							min: mapData[0],
							max: mapData[0] + (mapData[2] - 1),
						},
					}

					currentRanges = append(currentRanges, newRange)
				}
			}
		}
	}
	mappings[currentMapName] = currentRanges

	return data{
		seeds:    seeds,
		mappings: mappings,
	}
}

func getValue(ranges []mapRange, key int) int {
	for index := range ranges {
		current := ranges[index]
		if current.start.min <= key && current.start.max >= key {
			difference := key - current.start.min
			return current.destination.min + difference
		}
	}
	return key
}

func getRanges(ranges []mapRange, key rangeInfo) []rangeInfo {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start.min < ranges[j].start.min
	})

	result := make([]rangeInfo, 0)

	for index := range ranges {
		current := ranges[index]
		if key.min < current.start.min && key.max < current.start.min {
			result = append(result, key)
			return result
		}

		if key.min < current.start.min && key.max >= current.start.min && key.max <= current.start.max {
			result = append(result, rangeInfo{key.min, current.start.min - 1})
			difference := key.max - current.start.min
			result = append(result, rangeInfo{current.destination.min, current.destination.min + difference})
			return result
		}

		if key.min < current.start.min && key.max > current.start.max {
			result = append(result, rangeInfo{key.min, current.start.min - 1})
			result = append(result, rangeInfo{current.destination.min, current.destination.max})
			key = rangeInfo{current.start.max + 1, key.max}
			continue
		}

		if key.min >= current.start.min && key.max <= current.start.max {
			offsetStart := key.min - current.start.min
			difference := key.max - key.min
			result = append(result, rangeInfo{current.destination.min + offsetStart, current.destination.min + offsetStart + difference})
			return result
		}

		if key.min >= current.start.min && key.min <= current.start.max && key.max > current.start.max {
			offsetStart := key.min - current.start.min
			result = append(result, rangeInfo{current.destination.min + offsetStart, current.destination.max})
			key = rangeInfo{current.start.max + 1, key.max}
			continue
		}

		if key.min > current.start.max {
			continue
		}
	}

	result = append(result, key)

	return result
}

func Part1(filename string) int {
	minLocation := math.MaxInt

	d := readData(filename)
	for _, seed := range d.seeds {
		soil := getValue(d.mappings["seed-to-soil"], seed)
		fertilizer := getValue(d.mappings["soil-to-fertilizer"], soil)
		water := getValue(d.mappings["fertilizer-to-water"], fertilizer)
		light := getValue(d.mappings["water-to-light"], water)
		temperature := getValue(d.mappings["light-to-temperature"], light)
		humidity := getValue(d.mappings["temperature-to-humidity"], temperature)
		location := getValue(d.mappings["humidity-to-location"], humidity)
		
		minLocation = int(math.Min(float64(location), float64(minLocation)))
	}

	return minLocation
}

func Part2(filename string) int {
	minLocation := math.MaxInt

	d := readData(filename)

	seedRanges := make([]rangeInfo, 0)
	for i := 0; i < len(d.seeds); i += 2 {
		seedRanges = append(seedRanges, rangeInfo{
			min: d.seeds[i],
			max: d.seeds[i] + (d.seeds[i+1] - 1),
		})
	}

	soilRanges := make([]rangeInfo, 0)
	for _, seedRange := range seedRanges {
		soilRanges = append(soilRanges, getRanges(d.mappings["seed-to-soil"], seedRange)...)
	}

	fertilizerRanges := make([]rangeInfo, 0)
	for _, soilRange := range soilRanges {
		fertilizerRanges = append(fertilizerRanges, getRanges(d.mappings["soil-to-fertilizer"], soilRange)...)
	}

	waterRanges := make([]rangeInfo, 0)
	for _, fertilizerRange := range fertilizerRanges {
		waterRanges = append(waterRanges, getRanges(d.mappings["fertilizer-to-water"], fertilizerRange)...)
	}

	lightRanges := make([]rangeInfo, 0)
	for _, waterRange := range waterRanges {
		lightRanges = append(lightRanges, getRanges(d.mappings["water-to-light"], waterRange)...)
	}

	temperatureRanges := make([]rangeInfo, 0)
	for _, lightRange := range lightRanges {
		temperatureRanges = append(temperatureRanges, getRanges(d.mappings["light-to-temperature"], lightRange)...)
	}

	humidityRanges := make([]rangeInfo, 0)
	for _, temperatureRange := range temperatureRanges {
		humidityRanges = append(humidityRanges, getRanges(d.mappings["temperature-to-humidity"], temperatureRange)...)
	}

	locationRanges := make([]rangeInfo, 0)
	for _, humidityRange := range humidityRanges {
		locationRanges = append(locationRanges, getRanges(d.mappings["humidity-to-location"], humidityRange)...)
	}

	for i := range locationRanges {
		locationRange := locationRanges[i]
		minLocation = int(math.Min(float64(locationRange.min), float64(minLocation)))
	}

	return minLocation
}
