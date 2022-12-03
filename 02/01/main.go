package main

import (
	"bufio"
	"fmt"
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
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	beats = map[int]int{
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	}
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var score int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		opponent := choices[fields[0]]
		player := choices[fields[1]]

		if opponent == player {
			score += draw + player
			continue
		}

		if player == beats[opponent] {
			score += win + player
			continue
		}
		score += loss + player
	}
	fmt.Print(score)
}
