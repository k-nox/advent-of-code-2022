package partone

import (
	"bufio"
	"os"
	"strconv"
)

func Run() int {
	// open file
	file, err := os.Open("inputs/day01/input.txt")
	if err != nil {
		panic(err)
	}

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var curr int
	var max int

	// scan file one line at a time
	for scanner.Scan() {
		// get the current line
		currLine := scanner.Text()
		// if line is blank then we just finished a block
		if currLine == "" {
			if curr > max {
				max = curr
			}
			curr = 0
		} else {
			currNum, err := strconv.Atoi(currLine)
			if err != nil {
				panic(err)
			}
			curr += currNum
		}
	}

	return max
}
