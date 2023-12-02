package main

import (
	"fmt"
	"time"

	"github.com/gossie/advent-of-code-2023/day1"
	"github.com/gossie/advent-of-code-2023/day2"
)

func main() {

	fmt.Println("\nPerforming tasks of day1")
	startday1Part1 := time.Now().UnixMilli()
	day1Part1 := day1.Part1("day1/day1.txt")
	fmt.Println("day1, task 1: ", day1Part1, ", took", (time.Now().UnixMilli() - startday1Part1), "ms")
	startday1Part2 := time.Now().UnixMilli()
	day1Part2 := day1.Part2("day1/day1.txt")
	fmt.Println("day1, task 2: ", day1Part2, ", took", (time.Now().UnixMilli() - startday1Part2), "ms")

	fmt.Println("\nPerforming tasks of day2")
	startday2Part1 := time.Now().UnixMilli()
	day2Part1 := day2.Part1("day2/day2.txt")
	fmt.Println("day2, task 1: ", day2Part1, ", took", (time.Now().UnixMilli() - startday2Part1), "ms")
	startday2Part2 := time.Now().UnixMilli()
	day2Part2 := day2.Part2("day2/day2.txt")
	fmt.Println("day2, task 2: ", day2Part2, ", took", (time.Now().UnixMilli() - startday2Part2), "ms")
}
