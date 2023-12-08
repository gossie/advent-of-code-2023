package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	value       string
	left, right *node
}

func createOrGetNode(value string, allNodes map[string]*node) *node {
	if _, found := allNodes[value]; !found {
		allNodes[value] = &node{
			value: value,
		}
	}
	return allNodes[value]
}

func readData(filename string) (map[string]*node, []rune) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := make([]rune, 0)
	allNodes := make(map[string]*node)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if len(instructions) == 0 {
			for _, i := range line {
				instructions = append(instructions, i)
			}
		} else {
			nodeDeclaration := strings.Split(line, " = ")
			nodeValue := nodeDeclaration[0]
			children := strings.Split(nodeDeclaration[1][1:len(nodeDeclaration[1])-1], ", ")

			newNode := createOrGetNode(nodeValue, allNodes)
			newNode.left = createOrGetNode(children[0], allNodes)
			newNode.right = createOrGetNode(children[1], allNodes)
		}
	}

	return allNodes, instructions
}

func search(n *node, instructions []rune, targetCondition func(string) bool) int {
	steps := int(0)
	for {
		for _, i := range instructions {
			steps++
			switch i {
			case 'L':
				n = n.left
			case 'R':
				n = n.right
			default:
				panic(fmt.Sprintf("unknown instruction: %v", i))
			}

			if targetCondition(n.value) {
				return steps
			}
		}
	}
}

func leastCommonMultiple(differentSteps []int) int {
	currentSteps := differentSteps[0]
	for i := 1; i < len(differentSteps); i++ {
		if currentSteps == differentSteps[i] {
			continue
		}

		if currentSteps > differentSteps[i] {
			tmp := currentSteps
			for currentSteps%differentSteps[i] != 0 {
				currentSteps += tmp
			}
		} else {
			tmp := differentSteps[i]
			for differentSteps[i]%currentSteps != 0 {
				differentSteps[i] += tmp
			}
			currentSteps = differentSteps[i]
		}
	}
	return currentSteps
}

func Part1(filename string) int {
	allNodes, instructions := readData(filename)
	node := allNodes["AAA"]
	targetCondition := func(s string) bool {
		return s == "ZZZ"
	}

	return search(node, instructions, targetCondition)
}

func Part2(filename string) int {
	allNodes, instructions := readData(filename)

	targetCondition := func(s string) bool {
		return s[len(s)-1] == 'Z'
	}

	differentSteps := make([]int, 0)
	for _, n := range allNodes {
		if n.value[len(n.value)-1] == 'A' {
			differentSteps = append(differentSteps, search(n, instructions, targetCondition))
		}
	}

	return leastCommonMultiple(differentSteps)
}
