package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var points int
	for scanner.Scan() {
		currLine := scanner.Text()
		half := len(currLine) / 2
		set := mapset.NewSet[string]()
		repeated := mapset.NewSet[string]()

		for i, s := range currLine {
			if i < half {
				set.Add(string(s))
			} else {
				if set.Contains(string(s)) && !repeated.Contains(string(s)) {
					repeated.Add(string(s))
					if s >= 97 {
						points += int(s) - 96
					} else {
						points += int(s) - 38
					}
				}
			}
		}
	}
	fmt.Println(points)
}
