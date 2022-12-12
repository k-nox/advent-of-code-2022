package parttwo

import (
	"bufio"
	"strconv"
	"strings"
)

const (
	Blank  = "."
	Sprite = "#"
)

func Run(scanner *bufio.Scanner) string {
	cycle := 0
	middleOfSprite := 1
	output := strings.Builder{}

	for scanner.Scan() {
		instruction := strings.Fields(scanner.Text())
		// first thing -> start the cycle
		cycle++

		// what position are we in?
		output.WriteString(getPixel(cycle, middleOfSprite))

		if len(instruction) > 1 {
			// we are dealing with an addx instruction
			cycle++
			output.WriteString(getPixel(cycle, middleOfSprite))
			v, _ := strconv.Atoi(instruction[1])
			middleOfSprite += v
		}
	}
	return output.String()
}

func getPixel(cycle int, middleOfSprite int) string {
	pixel := Blank
	position := (cycle - 1) % 40
	if position == middleOfSprite || position == middleOfSprite-1 || position == middleOfSprite+1 {
		pixel = Sprite
	}
	if cycle%40 == 0 {
		pixel += "\n"
	}
	return pixel

}
