package parttwo

import (
	"bufio"
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"strconv"
	"strings"
)

type Monkey struct {
	Items *queue.Queue
	Operation Operation
	Test Test
	TestOperand int64
	TrueMonkey int
	FalseMonkey int
	InspectedCount int64
}

type Operation func(item int64) int64
type Test func(item int64) bool
type MonkeyCollection []*Monkey

func Run(scanner *bufio.Scanner) int64 {
	var (
		currRawMonkey []string
		rawMonkeys [][]string
	)

	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			rawMonkeys = append(rawMonkeys, currRawMonkey)
			currRawMonkey = nil
		} else {
			currRawMonkey = append(currRawMonkey, currLine)
		}
	}
	rawMonkeys = append(rawMonkeys, currRawMonkey)
	monkeys := BuildMonkeys(rawMonkeys)
	divisor := monkeys.GetDivisor()

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			// keep going for each item
			for !monkey.Items.Empty() {
				// get the item
				rawItem, err := monkey.Items.Get(1)
				if err != nil {
					panic(fmt.Errorf("error getting item from queue: %s", err.Error()))
				}
				// inspect the item
				item, ok := rawItem[0].(int64)
				if !ok {
					panic(fmt.Errorf("item isn't an int64 as expected: %v", item))
				}
				item = monkey.Operation(item)
				item = item % divisor
				if monkey.Test(item) {
					err = monkeys[monkey.TrueMonkey].Items.Put(item)
					if err != nil {
						panic(fmt.Errorf("error adding item %d to monkey %d queue: %s", item, monkey.TrueMonkey, err.Error()))
					}
				} else {
					err = monkeys[monkey.FalseMonkey].Items.Put(item)
					if err != nil {
						panic(fmt.Errorf("error adding item %d to monkey %d queue: %s", item, monkey.FalseMonkey, err.Error()))
					}
				}
				monkey.InspectedCount++
			}
		}
	}

	highestMonkeys := monkeys.GetIndexesOfTopTwoMonkeys()

	return monkeys[highestMonkeys[0]].InspectedCount * monkeys[highestMonkeys[1]].InspectedCount
}


func BuildMonkeys(rawMonkeys [][]string) MonkeyCollection {
	monkeys := MonkeyCollection{}
	for _, rawMonkey := range rawMonkeys {
		monkeys = append(monkeys, BuildMonkey(rawMonkey))
	}
	return monkeys
}

func BuildMonkey(rawMonkey []string) *Monkey {
	items := ParseItems(rawMonkey[1])
	operation := ParseOperation(rawMonkey[2])
	test, testOperand := ParseTest(rawMonkey[3])
	trueMonkey := ParseBoolMonkey(rawMonkey[4])
	falseMonkey := ParseBoolMonkey(rawMonkey[5])

	return &Monkey{
		Items: items,
		Operation: operation,
		Test: test,
		TestOperand: testOperand,
		TrueMonkey: trueMonkey,
		FalseMonkey: falseMonkey,
	}
}

func ParseItems(itemsLine string) *queue.Queue {
	itemsQueue := queue.New(3)
	itemsList := strings.Split(strings.Split(itemsLine, ":")[1], ",")
	for _, rawItem := range itemsList {
		item, err := strconv.ParseInt(strings.TrimSpace(rawItem), 10, 64)
		if err != nil {
			panic(fmt.Errorf("error parsing item %s: %s", rawItem, err.Error()))
		}
		err = itemsQueue.Put(item)
		if err != nil {
			panic(fmt.Errorf("error adding item %s to queue: %s", rawItem, err.Error()))
		}
	}
	return itemsQueue
}

func ParseOperation(operationLine string) Operation {
	// we only actually care about the last two fields in the list
	operationFields := strings.Fields(operationLine)[4:]
	operator := operationFields[0]
	operand, _ := strconv.ParseInt(operationFields[1], 10, 64)

	if operator == "+" {
		return func(item int64) int64 {
			if operand == 0 {
				operand = item
			}
			return item + operand
		}
	}

	return func(item int64) int64 {
		if operand == 0 {
			operand = item
		}
		return item * operand
	}
}

func ParseTest(testLine string) (Test, int64) {
	// we only care about the last field
	testList := strings.Fields(testLine)
	testNum, err := strconv.ParseInt(testList[len(testList) - 1], 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing testLine %s: %s", testLine, err.Error()))
	}
	return func(item int64) bool {
		return item % testNum == 0
	}, testNum
}

func ParseBoolMonkey(boolMonkey string) int {
	// we only care about the last field which is the monkey index
	boolMonkeyList := strings.Fields(boolMonkey)
	boolMonkeyIndex, err := strconv.Atoi(boolMonkeyList[len(boolMonkeyList) - 1])
	if err != nil {
		panic(fmt.Errorf("error parsing bool monkey line %s: %s", boolMonkey, err.Error()))
	}
	return boolMonkeyIndex
}

func (m MonkeyCollection) GetDivisor() int64 {
	var divisor int64 = 1
	for _, monkey := range m {
		divisor *= monkey.TestOperand
	}
	return divisor
}

func (m MonkeyCollection) GetIndexesOfTopTwoMonkeys() []int {
	largestA := 0
	largestB := 0

	for index, monkey := range m {
		if monkey.InspectedCount > m[largestA].InspectedCount {
			largestA = index
		} else if monkey.InspectedCount > m[largestB].InspectedCount {
			largestB = index
		}
	}
	return []int{largestA, largestB}
}