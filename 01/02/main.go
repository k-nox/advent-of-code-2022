package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open file
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var curr int
	maxes := []int{0, 0, 0}

	for scanner.Scan() {
		currLine := scanner.Text()
		// check if we reached the end of the block
		if currLine == "" {
			if curr > maxes[0] {
				maxes[2] = maxes[1]
				maxes[1] = maxes[0]
				maxes[0] = curr
			}
			// reset curr
			curr = 0
			// move on to the next line
			continue
		}
		// if we get here then we're within a block still
		currNum, err := strconv.Atoi(currLine)
		if err != nil {
			panic(err)
		}
		curr += currNum
	}

	// handle the last block
	if curr > maxes[0] {
		maxes[2] = maxes[1]
		maxes[1] = maxes[0]
		maxes[0] = curr
	}

	var total int

	for _, inidvMax := range maxes {
		total += inidvMax
	}

	fmt.Print(total)
}
