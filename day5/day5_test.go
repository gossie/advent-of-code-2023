package day5_test

import (
	"testing"

	"github.com/gossie/advent-of-code-2023/day5"
)

func TestPart1(t *testing.T) {
	part1 := day5.Part1("day5_test.txt")
	if part1 != 35 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day5.Part2("day5_test.txt")
	if part2 != 46 {
		t.Fatalf("part2 = %v", part2)
	}
}
