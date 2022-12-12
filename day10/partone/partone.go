package partone

import (
	"bufio"
	"strconv"
	"strings"
)

func Run(scanner *bufio.Scanner) int {
	target := 20
	incrementAmount := 40
	register := 1
	cycle := 0
	sum := 0
	for scanner.Scan() {
		instruction := strings.Fields(scanner.Text())
		// first thing -> start the cycle
		cycle++

		// check the signal strength
		if cycle == target {
			sum += (cycle * register)
			target += incrementAmount
		}

		if len(instruction) > 1 {
			// we are dealing with an addx instruction
			cycle++
			if cycle == target {
				sum += (cycle * register)
				target += incrementAmount
			}
			v, _ := strconv.Atoi(instruction[1])
			register += v
		}
	}
	return sum
}
