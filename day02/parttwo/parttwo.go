package parttwo

import (
	"bufio"
	"os"
	"strings"
)

const (
	rock int = iota + 1
	paper
	scissors

	loss = 0
	draw = 3
	win  = 6
)

var (
	choices = map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	endResults = map[string]int{
		"X": loss,
		"Y": draw,
		"Z": win,
	}

	beats = map[int]int{
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	}

	loses = map[int]int{
		rock:     scissors,
		paper:    rock,
		scissors: paper,
	}
)

func Run() int {
	f, err := os.Open("inputs/day02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var score int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		opponent := choices[fields[0]]
		desiredResult := endResults[fields[1]]
		var player int

		switch desiredResult {
		case loss:
			player = loses[opponent]
		case draw:
			player = opponent
		case win:
			player = beats[opponent]
		}

		score += desiredResult + player
	}
	return score
}
