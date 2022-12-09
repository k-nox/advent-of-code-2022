package partone

import (
	"bufio"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

func Run() int {
	// goal: given a string of letters, find the index + 1 of the end of the first substring of 4 unique characters
	f, err := os.Open("inputs/day06/input.txt")
	if err != nil {
		panic(err)
	}
	var startOfPacket int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		stream := scanner.Text()
		// loop through stream starting at the 4th char
		for i := 3; i < len(stream); i++ {
			// check the 4 char window from i - 3 to i
			window := stream[i-3 : i+1]
			isWindowPacket := checkIsWindowPacket(window)
			if isWindowPacket {
				startOfPacket = i + 1
				break
			}
		}
	}
	return startOfPacket

}

func checkIsWindowPacket(window string) bool {
	uniqueSet := mapset.NewSet[rune]()
	for _, char := range window {
		if uniqueSet.Contains(char) {
			return false
		} else {
			uniqueSet.Add(char)
		}
	}
	return true
}
