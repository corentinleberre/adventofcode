package day7

import (
	"adventofcode/src/util"
	"sort"
	"strconv"
	"strings"
)

func buildTree(dataset []string) *Directory {
	tree := &Directory{
		parentDirectory: nil,
		name:            "/",
	}
	var currentDirectory *Directory
	for _, line := range dataset {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "$":
			if parts[1] == "cd" {
				switch parts[2] {
				case "/":
					currentDirectory = tree
				case "..":
					currentDirectory = currentDirectory.parentDirectory
				default:
					var directory *Directory
					for _, child := range currentDirectory.directories {
						if child.name == parts[2] {
							directory = child
						}
					}
					if directory == nil {
						directory = &Directory{
							parentDirectory: currentDirectory,
							name:            parts[2],
						}
						currentDirectory.directories = append(currentDirectory.directories, directory)
					}
					currentDirectory = directory
				}
			}
		case "dir":
			dir := &Directory{
				parentDirectory: currentDirectory,
				name:            parts[1],
			}
			if currentDirectory.directories == nil {
				currentDirectory.directories = []*Directory{}
			}
			currentDirectory.directories = append(currentDirectory.directories, dir)

		default:
			value, _ := strconv.Atoi(parts[0])
			file := File{parentDirectory: currentDirectory, size: int64(value), name: parts[1]}
			if currentDirectory.files == nil {
				currentDirectory.files = []File{}
			}
			currentDirectory.files = append(currentDirectory.files, file)
		}
	}
	return tree
}

func calcSize(dir *Directory, directoryAtMost100k *int64, directoriesToDeleteToUpdateSystem *[]int64) {
	if dir == nil {
		return
	}
	if len(dir.files) > 0 {
		var totalSize int64
		for _, file := range dir.files {
			totalSize += file.size
		}
		dir.totalSize = totalSize
	}
	if dir.directories != nil {
		for _, subdir := range dir.directories {
			calcSize(subdir, directoryAtMost100k, directoriesToDeleteToUpdateSystem)
			dir.totalSize += subdir.totalSize
		}
	}
	if dir.totalSize <= 100000 {
		*directoryAtMost100k += dir.totalSize
	}
	if dir.totalSize >= (30000000 - 23909866) {
		*directoriesToDeleteToUpdateSystem = append(*directoriesToDeleteToUpdateSystem, dir.totalSize)
	}
}

func ChallengeOfTheDay(dataset []string) util.ChallengeResult {
	tree := buildTree(dataset)

	var directoryAtMost100k int64
	var directoriesToDeleteToUpdateSystem []int64

	calcSize(tree, &directoryAtMost100k, &directoriesToDeleteToUpdateSystem)

	sort.Slice(directoriesToDeleteToUpdateSystem, func(i, j int) bool {
		return directoriesToDeleteToUpdateSystem[i] < directoriesToDeleteToUpdateSystem[j]
	})

	return util.ChallengeResult{Puzzle1: int(directoryAtMost100k), Puzzle2: int(directoriesToDeleteToUpdateSystem[0])}
}
