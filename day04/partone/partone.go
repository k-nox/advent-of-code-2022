package partone

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Run() int {
	// input: [[n-m, k-i], ...]
	// find cases where k >= n && i <= m || n >= k && m <= i
	// meaning: second pair should have a start that is greater than or equal to first pair start, and end that is less than or equal to second pair end, or vice versa
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
		// check if first is contained within second
		if (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd) {
			points += 1
		}
	}
	return points
}
