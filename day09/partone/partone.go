package partone

import (
	"bufio"
	"image"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	Right      = image.Pt(1, 0)
	Left       = image.Pt(-1, 0)
	Up         = image.Pt(0, 1)
	Down       = image.Pt(0, -1)
	directions = map[string]image.Point{
		"R": Right,
		"U": Up,
		"L": Left,
		"D": Down,
	}
)

type Positions struct {
	Head image.Point
	Tail image.Point
}

func Run(scanner *bufio.Scanner) int {
	positions := &Positions{}

	tailVisited := mapset.NewSet[image.Point]()
	tailVisited.Add(positions.Tail)

	for scanner.Scan() {
		direction, count := ParseInstructions(scanner.Text())
		for i := 0; i < count; i++ {
			positions.ApplyInstruction(direction)
			tailVisited.Add(positions.Tail)
		}

	}
	return tailVisited.Cardinality()
}

func (p *Positions) ApplyInstruction(direction image.Point) {
	// move head
	p.Head = p.Head.Add(direction)

	// check if tail needs to be updated
	diff := p.Head.Sub(p.Tail)
	if diff.X <= 1 && diff.Y <= 1 && diff.X >= -1 && diff.Y >= -1 {
		// tail is touching either diagonally or laterally
		return
	}

	var tailTranslation image.Point
	if diff.Y > 0 {
		tailTranslation = tailTranslation.Add(Up)
	}
	if diff.Y < 0 {
		tailTranslation = tailTranslation.Add(Down)
	}
	if diff.X < 0 {
		tailTranslation = tailTranslation.Add(Left)
	}
	if diff.X > 0 {
		tailTranslation = tailTranslation.Add(Right)
	}

	p.Tail = p.Tail.Add(tailTranslation)

}

func ParseInstructions(rawInstructions string) (image.Point, int) {
	instructions := strings.Fields(rawInstructions)
	count, err := strconv.Atoi(instructions[1])
	if err != nil {
		panic(err)
	}
	return directions[instructions[0]], count
}
