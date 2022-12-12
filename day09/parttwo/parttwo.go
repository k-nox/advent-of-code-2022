package parttwo

import (
	"bufio"
	"image"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/k-nox/advent-of-code-2022/day09/partone"
)

func Run(scanner *bufio.Scanner) int {
	positions := make([]image.Point, 10)
	tail := len(positions) - 1
	tailVisited := mapset.NewSet[image.Point]()
	tailVisited.Add(positions[tail])

	for scanner.Scan() {
		direction, count := partone.ParseInstructions(scanner.Text())
		for i := 0; i < count; i++ {
			positions = ApplyInstruction(direction, positions)
			tailVisited.Add(positions[tail])
		}
	}

	return tailVisited.Cardinality()
}

func ApplyInstruction(direction image.Point, positions []image.Point) []image.Point {
	// move head
	positions[0] = positions[0].Add(direction)
	UpdatePositions(positions, 0, 1)
	return positions
}

func UpdatePositions(positions []image.Point, prev int, curr int) {
	if curr >= len(positions) {
		// we're all done
		return
	}
	diff := positions[prev].Sub(positions[curr])
	if diff.X <= 1 && diff.Y <= 1 && diff.X >= -1 && diff.Y >= -1 {
		// curr is already touching prev either diagonally or laterally
		return
	}

	var translation image.Point
	if diff.Y > 0 {
		translation = translation.Add(partone.Up)
	}
	if diff.Y < 0 {
		translation = translation.Add(partone.Down)
	}
	if diff.X < 0 {
		translation = translation.Add(partone.Left)
	}
	if diff.X > 0 {
		translation = translation.Add(partone.Right)
	}

	positions[curr] = positions[curr].Add(translation)
	UpdatePositions(positions, curr, curr+1)
}
