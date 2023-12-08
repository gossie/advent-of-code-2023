package day8_test

import (
	"testing"

	"github.com/gossie/advent-of-code-2023/day8"
)

func TestPart1(t *testing.T) {
	part1 := day8.Part1("day8_test1.txt")
	if part1 != 2 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day8.Part2("day8_test2.txt")
	if part2 != 6 {
		t.Fatalf("part2 = %v", part2)
	}
}
