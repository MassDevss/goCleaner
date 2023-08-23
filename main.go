package main

import (
	f "fmt"
	"io/fs"
	"os"

	h "github.com/ManuelsSaNt/goCleaner/helpers"
)

var dirToSwap string = ""
var deppest int = 0

func main() {

	dirToClean := ".."
	dirToSwap = dirToClean

	// var toDelete []string = []string{"node_modules"}
	// baseDir := readADir(dirToClean)

	// recursiveCleaner(baseDir, toDelete)
	recursiveNavigator(dirToClean)

	// f.Println()

}

func recursiveNavigator(setterPath string) {

	baseDirCollection := readADir(setterPath)
	var directories = []fs.DirEntry{}

	for _, childDir := range baseDirCollection {

		if childDir.IsDir() && childDir.Name() != "node_modules" {
			directories = append(directories, childDir)
		}

		f.Println(repeat(" ", deppest) + childDir.Name())
	}

	// deppest++

	for _, parent := range directories {

		parentPath := setterPath + "/" + parent.Name()

		f.Println("direcotry: " + parent.Name())
		f.Println("path: " + parentPath)

		recursiveNavigator(parentPath)
	}

}

func readADir(path string) []fs.DirEntry {
	dir, err := os.ReadDir(path)
	h.ManageErr(err)

	return dir
}

func printDir(path string) {
	dir, err := os.ReadDir(path)
	h.ManageErr(err)

	for _, dirOrFile := range dir {
		f.Println(dirOrFile)
	}
}

func repeat(toRepeat string, times int) string {
	newTxt := ""

	for i := 0; i < times; i++ {
		newTxt += toRepeat
	}

	return newTxt
}
