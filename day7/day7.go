package day7

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var pointsByCard map[rune]int = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func iterateCards(cards []int, winFunc func(int, int) bool, resetFunc func()) bool {
	for jokerValue := 2; jokerValue <= 14; jokerValue++ {
		resetFunc()

		matchingIndizes := make([]int, 0)
		for i := 0; i < len(cards); i++ {
			currentCard := cards[i]
			if currentCard == 1 {
				currentCard = jokerValue
			}
			hits := 0
			for j := i + 1; j < len(cards); j++ {
				if slices.Contains(matchingIndizes, j) {
					continue
				}

				otherCard := cards[j]
				if otherCard == 1 {
					otherCard = jokerValue
				}

				if currentCard == otherCard {
					hits++
					matchingIndizes = append(matchingIndizes, j)
				}
			}

			if winFunc(hits, currentCard) {
				return true
			}
		}
	}
	return false
}

func isFiveOfAKind(cards []int) bool {
	winFunc := func(hits int, currentCard int) bool {
		return hits == 4
	}

	resetFunc := func() {}

	return iterateCards(cards, winFunc, resetFunc)
}

func isFourOfAKind(cards []int) bool {
	winFunc := func(hits int, currentCard int) bool {
		return hits == 3
	}

	resetFunc := func() {}

	return iterateCards(cards, winFunc, resetFunc)
}

func isFullHouse(cards []int) bool {
	foundPair := false
	foundTriplets := false

	winFunc := func(hits int, currentCard int) bool {
		if hits == 2 {
			foundTriplets = true
		}

		if hits == 1 {
			foundPair = true
		}

		return foundTriplets && foundPair
	}

	resetFunc := func() {
		foundPair = false
		foundTriplets = false
	}

	return iterateCards(cards, winFunc, resetFunc)
}

func isThreeOfAKind(cards []int) bool {
	winFunc := func(hits int, currentCard int) bool {
		return hits == 2
	}

	resetFunc := func() {}

	return iterateCards(cards, winFunc, resetFunc)
}

func isTwoPairs(cards []int) bool {
	firstPair := 0
	secondPair := 0

	winFunc := func(hits int, currentCard int) bool {
		if hits == 1 && firstPair == 0 {
			firstPair = currentCard
		}

		if hits == 1 && firstPair > 0 && firstPair != currentCard {
			secondPair = currentCard
		}

		return firstPair > 0 && secondPair > 0 && firstPair != secondPair
	}

	resetFunc := func() {
		firstPair = 0
		secondPair = 0
	}

	return iterateCards(cards, winFunc, resetFunc)
}

func isOnePair(cards []int) bool {
	winFunc := func(hits int, currentCard int) bool {
		return hits == 1
	}

	resetFunc := func() {}

	return iterateCards(cards, winFunc, resetFunc)
}

type hand struct {
	cards []int
	bid   int
}

func (h *hand) handType() int {
	if isFiveOfAKind(h.cards) {
		return fiveOfAKind
	}

	if isFourOfAKind(h.cards) {
		return fourOfAKind
	}

	if isFullHouse(h.cards) {
		return fullHouse
	}

	if isThreeOfAKind(h.cards) {
		return threeOfAKind
	}

	if isTwoPairs(h.cards) {
		return twoPairs
	}

	if isOnePair(h.cards) {
		return onePair
	}

	return highCard
}

func readData(filename string, withJoker bool) []hand {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hands := make([]hand, 0)

	for scanner.Scan() {
		line := scanner.Text()
		newHand := hand{}
		handStr := strings.Split(line, " ")
		for _, c := range handStr[0] {
			value := pointsByCard[c]
			if withJoker && c == 'J' {
				value = 1
			}
			newHand.cards = append(newHand.cards, value)
		}

		bid, _ := strconv.Atoi(handStr[1])
		newHand.bid = bid

		hands = append(hands, newHand)
	}

	return hands
}

func perform(hands []hand) int {
	sort.Slice(hands, func(i, j int) bool {
		typeI := hands[i].handType()
		typeJ := hands[j].handType()
		if typeI != typeJ {
			return typeI < typeJ
		}

		for c := 0; c < len(hands[i].cards); c++ {
			if hands[i].cards[c] != hands[j].cards[c] {
				return hands[i].cards[c] < hands[j].cards[c]
			}
		}
		return false
	})

	totalWinnings := 0
	for i := range hands {
		totalWinnings += hands[i].bid * (i + 1)
	}

	return totalWinnings
}

func Part1(filename string) int {
	hands := readData(filename, false)
	return perform(hands)
}

func Part2(filename string) int {
	hands := readData(filename, true)
	return perform(hands)
}
