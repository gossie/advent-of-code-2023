package main

import (
    "fmt"
    "time"
    "github.com/gossie/adventofcode2023/day1"
)

func main() {
    

    fmt.Println("\nPerforming tasks of day1")
    startday1Part1 := time.Now().UnixMilli()
    day1Part1 := day1.Part1("day1/day1.txt")
    fmt.Println("day1, task 1: ", day1Part1, ", took", (time.Now().UnixMilli() - startday1Part1), "ms")
    startday1Part2 := time.Now().UnixMilli()
    day1Part2 := day1.Part2("day1/day1.txt")
    fmt.Println("day1, task 2: ", day1Part2, ", took", (time.Now().UnixMilli() - startday1Part2), "ms")
}
