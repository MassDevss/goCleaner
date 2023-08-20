package main

import (
	f "fmt"
	"io/fs"
	"os"

	h "github.com/ManuelsSaNt/goCleaner/helpers"
)

func main() {

	dirToClean := "../"
	var toDelete []string = []string{"node_modules"}
	RootDir := readADir(dirToClean)

}

func recursiveCleaner(rootDir []fs.DirEntry, toDelete []string) {

	extraDirs := []fs.DirEntry{}

	for _, dirOrFile := range rootDir {

		isTarget := false

		// for _,

		if dirOrFile.IsDir() {
			extraDirs = append(extraDirs, dirOrFile)
		}
	}

}

func readADir(path string) []fs.DirEntry {
	dir, err := os.ReadDir("../")
	h.ManageErr(err)

	return dir
}

func printDir(path string) {
	dir, err := os.ReadDir("../")
	h.ManageErr(err)

	for _, dirOrFile := range dir {
		f.Println(dirOrFile)
	}
}
