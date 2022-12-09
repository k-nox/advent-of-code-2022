package parttwo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Run() string {
	f, err := os.Open("inputs/day05/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	stackLines := []string{}
	parsingStack := true
	var stacks [][]string

	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			if parsingStack {
				parsingStack = false
				stacks = parseStacks(stackLines)
			}
			continue
		}
		if parsingStack {
			stackLines = append(stackLines, currLine)
		} else {
			stacks = applyInstructions(stacks, currLine)
		}
	}
	tops := getTopItems(stacks)
	return tops
}

func getTopItems(stacks [][]string) string {
	topItems := ""
	for _, stack := range stacks {
		topItems += stack[len(stack)-1]
	}
	return topItems
}

func applyInstructions(stacks [][]string, rawInstruction string) [][]string {
	instructions := strings.Fields(rawInstruction)
	count, _ := strconv.Atoi(instructions[1])
	startStack, _ := strconv.Atoi(instructions[3])
	startStack--
	endStack, _ := strconv.Atoi(instructions[5])
	endStack--
	groupStartIndex := len(stacks[startStack]) - count
	items := stacks[startStack][groupStartIndex:]
	stacks[startStack] = stacks[startStack][:groupStartIndex]
	stacks[endStack] = append(stacks[endStack], items...)
	return stacks
}

func parseStacks(rawStacks []string) [][]string {
	rawHeader := rawStacks[len(rawStacks)-1]
	parsedHeader := strings.Fields(rawHeader)
	stacks := make([][]string, len(parsedHeader))

	for i := len(rawStacks) - 2; i >= 0; i-- {
		row := rawStacks[i]
		j := 0
		for k := 1; k < len(row); k += 4 {
			if curr := string(row[k]); curr != " " {
				stacks[j] = append(stacks[j], string(row[k]))
			}
			j++
		}
	}
	return stacks
}
