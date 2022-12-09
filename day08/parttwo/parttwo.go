package parttwo

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

	score := findBestScenicScore(grid)
	return score

}

func findBestScenicScore(grid [][]int) int {
	// look at each tree
	var bestScore int
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			score := findScenicScore(grid, row, col)
			if bestScore == 0 || score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func findScenicScore(grid [][]int, row int, col int) int {
	var (
		scoreAbove int
		scoreBelow int
		scoreLeft  int
		scoreRight int
	)

	tree := grid[row][col]

	// check left
	for i := col - 1; i >= 0; i-- {
		scoreLeft++
		if grid[row][i] >= tree {
			break
		}
	}

	// check right
	for i := col + 1; i < len(grid[row]); i++ {
		scoreRight++
		if grid[row][i] >= tree {
			break
		}
	}

	// check above
	for i := row - 1; i >= 0; i-- {
		scoreAbove++
		if grid[i][col] >= tree {
			break
		}
	}

	// check below
	for i := row + 1; i < len(grid); i++ {
		scoreBelow++
		if grid[i][col] >= tree {
			break
		}
	}

	return scoreAbove * scoreBelow * scoreRight * scoreLeft
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
