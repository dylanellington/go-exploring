package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LeetcodeProblemDirectory struct {
	path string
	LeetcodeProblem
}

func (dir *LeetcodeProblemDirectory) CreateReadme(chromium *Chromium) error {
	return dir.LeetcodeProblem.CreateReadme(chromium, dir.path)
}

func main() {
	chromium := Chromium{}
	defer chromium.Close()
	chromium.Start()

	leetcodeProblemDirectories := GetLeetcodeProblemDirectories()

	for _, problemDirectory := range leetcodeProblemDirectories {
		if !fileExists(problemDirectory.path + "/README.md") {
			fmt.Printf("Creating readme for leetcode problem %d at directory %s.\n", problemDirectory.number, problemDirectory.path)

			if err := problemDirectory.CreateReadme(&chromium); err != nil {
				panic(err)
			}
		}
	}
}

func GetLeetcodeProblemDirectories() []LeetcodeProblemDirectory {
	problemDirectories := make([]LeetcodeProblemDirectory, 0, 0)

	err := filepath.Walk("leetcode_problems", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			folderProblemNumber := strings.Split(info.Name(), "_")[0]

			if problemNumber, err := strconv.ParseInt(folderProblemNumber, 10, 32); err == nil {
				leetcodeProblem, _ := NewLeetcodeProblem(int(problemNumber))

				problemDirectory := LeetcodeProblemDirectory {
					path: path,
					LeetcodeProblem: leetcodeProblem,
				}

				problemDirectories = append(problemDirectories, problemDirectory)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return problemDirectories
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)

	if os.IsNotExist(err) || info.IsDir() {
		return false
	}

	return true
}
