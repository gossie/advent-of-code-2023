package day4_test

import (
	"testing"

	"github.com/gossie/advent-of-code-2023/day4"
)

func TestPart1(t *testing.T) {
	part1 := day4.Part1("day4_test.txt")
	if part1 != 13 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day4.Part2("day4_test.txt")
	if part2 != 30 {
		t.Fatalf("part2 = %v", part2)
	}
}
