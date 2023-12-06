package day6_test

import (
	"testing"

	"github.com/gossie/advent-of-code-2023/day6"
)

func TestPart1(t *testing.T) {
	part1 := day6.Part1("day6_test.txt")
	if part1 != 288 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day6.Part2("day6_test.txt")
	if part2 != 71503 {
		t.Fatalf("part2 = %v", part2)
	}
}
