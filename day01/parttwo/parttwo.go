package parttwo

import (
	"bufio"
	"errors"
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

	nLargest, err := findNLargest(scanner, 3)
	if err != nil {
		panic(err)
	}

	var total int

	for _, inidvLargest := range nLargest {
		total += inidvLargest
	}

	return total
}

func findNLargest(scanner *bufio.Scanner, n int) ([]int, error) {
	var curr int
	nLargest := make([]int, n)
	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			if curr > nLargest[0] {
				updateNLargest(&nLargest, 0, curr)
			}
			curr = 0
			continue
		}
		currNum, err := strconv.Atoi(currLine)
		if err != nil {
			return nil, err
		}
		curr += currNum
	}

	if curr > nLargest[0] {
		updateNLargest(&nLargest, 0, curr)
	}

	return nLargest, nil

}

func updateNLargest(nLargest *[]int, index int, newVal int) error {
	if index == len(*nLargest) {
		return nil
	}
	if index < 0 {
		return errors.New("index must be greater than 0")
	}
	if index > len(*nLargest) {
		return errors.New("index must be less thatn length of m")
	}
	nextVal := (*nLargest)[index]
	(*nLargest)[index] = newVal
	updateNLargest(nLargest, index+1, nextVal)
	return nil
}
