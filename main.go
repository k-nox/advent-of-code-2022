package main

import (
	"bufio"
	"fmt"
	"github.com/k-nox/advent-of-code-2022/day09/partone"
	"os"
)

func main() {
	scanner := getScanner("09", false)
	result := partone.Run(scanner)
	fmt.Println(result)
}

func getScanner(day string, isTest bool) *bufio.Scanner {
	fileName := "input.txt"
	if isTest {
		fileName = "test.txt"
	}
	fileFullPath := fmt.Sprintf("inputs/day%s/%s", day, fileName)
	f, err := os.Open(fileFullPath)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(f)

}