package main

import (
	"bufio"
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	// goal: given a string of letters, find the index + 1 of the end of the first substring of 4 unique characters
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var startOfPacket int
		stream := scanner.Text()
		// loop through stream starting at the 4th char
		for i := 13; i < len(stream); i++ {
			// check the 4 char window from i - 3 to i
			window := stream[i-13 : i+1]
			isWindowPacket := checkIsWindowPacket(window)
			if isWindowPacket {
				startOfPacket = i + 1
				break
			}
		}
		fmt.Println(startOfPacket)
	}
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
