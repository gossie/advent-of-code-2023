package day1_test

import (
	"testing"

	"github.com/gossie/advent-of-code-2023/day1"
)

func TestPart1(t *testing.T) {
	part1 := day1.Part1("day1_test1.txt")
	if part1 != 142 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day1.Part2("day1_test2.txt")
	if part2 != 281 {
		t.Fatalf("part2 = %v", part2)
	}
}
