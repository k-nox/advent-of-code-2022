package partone

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	parent      *directory
	name        string
	children    []*directory
	files       []*file
	sizeOfFiles int
}

type file struct {
	name string
	size int
}

const (
	Root     = "/"
	CmdStart = "$"
	Dir      = "dir"
	Change   = "cd"
	List     = "ls"
	Previous = ".."
)

func Run() int {
	f, err := os.Open("inputs/day07/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var currentDir *directory
	var rootDir *directory
	listMode := false

	// parse and build directory tree
	for scanner.Scan() {
		currLine := scanner.Text()
		fields := strings.Fields(currLine)

		if fields[0] == CmdStart {
			// we are in a cmd
			listMode, currentDir = handleCmd(fields, currentDir)
			// always know where the root is
			if rootDir == nil {
				rootDir = currentDir
			}
		} else if listMode {
			// time to start adding files and directories
			processDirOrFile(currentDir, fields)
		}

	}
	sum := findSumOfAllDesirableDirs(rootDir)
	return sum

}

func findSumOfAllDesirableDirs(currentDir *directory) int {
	var sum int
	currSize := calculateDirSize(currentDir)
	if currSize <= 100000 {
		sum += currSize
	}
	for _, child := range currentDir.children {
		sum += findSumOfAllDesirableDirs(child)
	}
	return sum
}

func calculateDirSize(currentDir *directory) int {
	var size int
	// first sum up the sizes of the files if it's not already known
	if currentDir.sizeOfFiles == 0 && len(currentDir.files) > 0 {
		for _, file := range currentDir.files {
			currentDir.sizeOfFiles += file.size
		}
	}
	// now find sum of all the subdirectory files
	for _, dir := range currentDir.children {
		size += calculateDirSize(dir)
	}
	return size + currentDir.sizeOfFiles
}

func processDirOrFile(currentDir *directory, fields []string) {
	// ----- directories -----
	if fields[0] == Dir {
		currentDir.children = append(currentDir.children, createNewDir(fields[1], currentDir))
	} else {
		currentDir.files = append(currentDir.files, createNewFile(fields[0], fields[1]))
	}
}

func handleCmd(fields []string, currentDir *directory) (bool, *directory) {
	newDir := currentDir
	listMode := false

	// ------- CD ------
	// we are navigating
	if fields[1] == Change {
		newDirName := fields[2]
		newDir = cd(currentDir, newDirName)
	}

	// ------- LS -------
	// put us into list mode
	if fields[1] == List {
		listMode = true
	}

	return listMode, newDir

}

func cd(currentDir *directory, newDirName string) *directory {
	var newDir *directory
	// we're at the root, craete the first dir
	if currentDir == nil && newDirName == Root {
		newDir = createNewDir(newDirName, nil)
	} else if newDirName == Previous {
		// naviate back up the tree
		newDir = currentDir.parent
	} else {
		// update current dir
		for _, dir := range currentDir.children {
			if dir.name == newDirName {
				newDir = dir
			}
		}
	}
	return newDir
}

func createNewDir(dirName string, parent *directory) *directory {
	return &directory{
		parent:   parent,
		name:     dirName,
		children: []*directory{},
		files:    []*file{},
	}
}

func createNewFile(rawSize string, fileName string) *file {
	size, err := strconv.Atoi(rawSize)
	if err != nil {
		panic(err)
	}
	return &file{
		name: fileName,
		size: size,
	}
}
