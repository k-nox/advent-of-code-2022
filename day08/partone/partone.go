package partone

import (
	"bufio"
	"os"
	"strconv"
)

func Run() int {
	f, err := os.Open("inputs/day08/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	grid := parseInput(scanner)

	// find number of trees around the edges
	// assume every row is the same length
	edges := (len(grid) * 2) + (len(grid[0]) * 2) - 4

	innerVisible := findInnerVisibleTrees(grid)

	visible := edges + innerVisible
	return visible

}

func max(arr []int) int {
	m := 0
	for _, n := range arr {
		if m == 0 || n > m {
			m = n
		}
	}
	return m
}

func findInnerVisibleTrees(grid [][]int) int {
	count := 0
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			tree := grid[row][col]

			// check the right
			if tree > max(grid[row][col+1:]) {
				count++
				continue
			}

			// check the left
			if tree > max(grid[row][:col]) {
				count++
				continue
			}

			// check columns
			colAbove := []int{}
			colBelow := []int{}
			for i := 0; i < len(grid); i++ {
				if i < row {
					colAbove = append(colAbove, grid[i][col])
				}
				if i > row {
					colBelow = append(colBelow, grid[i][col])
				}
			}
			if tree > max(colAbove) {
				count++
				continue
			}
			if tree > max(colBelow) {
				count++
				continue
			}
		}
	}
	return count
}

func parseInput(scanner *bufio.Scanner) [][]int {
	grid := [][]int{}

	for scanner.Scan() {
		currLine := scanner.Text()
		// crate the row
		row := []int{}
		for _, tree := range currLine {
			tree, err := strconv.Atoi(string(tree))
			if err != nil {
				panic(err)
			}
			row = append(row, tree)
		}
		// add the row to the matrix
		grid = append(grid, row)
	}

	return grid
}
