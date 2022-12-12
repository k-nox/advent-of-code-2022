package partone

import (
	"bufio"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int
	operation      Operation
	test           Test
	inspectedCount int
}

type Operation struct {
	opType  string
	operand int
}

type Test struct {
	operand int
	ifTrue  int
	ifFalse int
}

const (
	Add = "+"
)

func Run(scanner *bufio.Scanner) int {

	rawMonkey := []string{}
	monkeys := []*Monkey{}

	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			monkeys = append(monkeys, BuildMonkey(rawMonkey))
			rawMonkey = nil
		} else {
			rawMonkey = append(rawMonkey, currLine)
		}
	}

	monkeys = append(monkeys, BuildMonkey(rawMonkey))

	for round := 0; round < 20; round++ {
		RunRound(monkeys)
	}

	// find the most active monkeys
	topTwo := make([]int, 2)
	for _, monkey := range monkeys {
		if monkey.inspectedCount > topTwo[0] {
			topTwo[1] = topTwo[0]
			topTwo[0] = monkey.inspectedCount
		}
	}

	return topTwo[0] * topTwo[1]
}

func RunRound(monkeys []*Monkey) {

	for _, monkey := range monkeys {
		monkey.RunTurn(monkeys)
	}
}

func (m *Monkey) RunTurn(monkeys []*Monkey) {
	for _, item := range m.items {
		itemToThrow := m.InspectItem(item)
		itemToThrow /= 3
		monkeyToThrowTo := m.TestItem(itemToThrow)
		monkeys[monkeyToThrowTo].items = append(monkeys[monkeyToThrowTo].items, itemToThrow)
		m.inspectedCount++
	}
	m.items = nil
}

func (m *Monkey) TestItem(item int) int {
	if item%m.test.operand == 0 {
		return m.test.ifTrue
	}
	return m.test.ifFalse
}

func (m *Monkey) InspectItem(item int) int {
	operand := m.operation.operand
	if operand == 0 {
		// hack because some monkeys multiply the item by itself
		operand = item
	}
	if m.operation.opType == Add {
		return item + operand
	}
	return item * operand
}

func BuildMonkey(rawMonkey []string) *Monkey {
	items := ParseItems(rawMonkey[1])
	operation := ParseOperation(rawMonkey[2])
	test := ParseTest(rawMonkey[3:])

	return &Monkey{
		items:          items,
		operation:      operation,
		test:           test,
		inspectedCount: 0,
	}
}

func ParseItems(rawItems string) []int {
	var items []int
	// grab just the items
	rawItems = strings.Split(rawItems, ":")[1]
	rawItemsList := strings.Split(rawItems, ",")

	for _, rawItem := range rawItemsList {
		item, _ := strconv.Atoi(strings.TrimSpace(rawItem))
		items = append(items, item)
	}
	return items
}

func ParseOperation(rawOperation string) Operation {
	rawOps := strings.Fields(rawOperation)
	operand, _ := strconv.Atoi(rawOps[5])
	return Operation{
		opType:  rawOps[4],
		operand: operand,
	}
}

func ParseTest(rawTest []string) Test {
	operand, _ := strconv.Atoi(strings.Fields(rawTest[0])[3])
	ifTrue, _ := strconv.Atoi(strings.Fields(rawTest[1])[5])
	ifFalse, _ := strconv.Atoi(strings.Fields(rawTest[2])[5])
	return Test{
		operand: operand,
		ifTrue:  ifTrue,
		ifFalse: ifFalse,
	}

}
