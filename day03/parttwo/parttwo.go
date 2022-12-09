package parttwo

import (
	"bufio"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

func Run() int {
	f, err := os.Open("inputs/day03/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var groupPos int
	var points int
	sets := []mapset.Set[string]{mapset.NewSet[string](), mapset.NewSet[string]()}

	// read through one line at a time
	for scanner.Scan() {
		// grab the text of the current line
		currLine := scanner.Text()
		// are we on the first or second elf of the group?
		if groupPos < 2 {
			for _, s := range currLine {
				sets[groupPos].Add(string(s))
			}
			groupPos++
		} else {
			// we are on the last elf of the group
			intersection := sets[0].Intersect(sets[1])
			for _, s := range currLine {
				if intersection.Contains(string(s)) {
					if s >= 97 {
						points += int(s) - 96
					} else {
						points += int(s) - 38
					}
					break
				}
			}
			groupPos = 0
			sets = []mapset.Set[string]{mapset.NewSet[string](), mapset.NewSet[string]()}
		}
	}
	return points
}
