package parttwo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Run() int {
	file, err := os.Open("inputs/day04/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var points int

	for scanner.Scan() {
		currLine := scanner.Text()
		pair := strings.Split(currLine, ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")
		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])

		if secondStart > firstStart && secondStart <= firstEnd {
			points++
			continue
		}

		if firstStart > secondStart && firstStart <= secondEnd {
			points++
			continue
		}

		if firstStart == secondStart || firstEnd == secondEnd {
			points++
			continue
		}

		// check if first is contained within second
		//if (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd) {
		//	points += 1
		//}
	}
	return points
}
