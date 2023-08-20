package main

import (
	f "fmt"
	"io/fs"
	"os"

	h "github.com/ManuelsSaNt/goCleaner/helpers"
)

func main() {

	// var toDelete string = "node_modules"

	printDir("../")

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
