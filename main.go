package main

import (
	"flag"
	f "fmt"
	"io/fs"
	"os"
	"strings"

	h "github.com/ManuelsSaNt/goCleaner/helpers"
	"golang.org/x/exp/slices"
)

var dirToSwap string = ""

func main() {

	helpFlag := flag.Bool("help", false, "Help comand use to show information")
	// ! config for flags
	rootSearch := flag.String("root", "", "the initial dir for search the delete files.")
	deleteFileOrDir := flag.String("target", "", "The files to delete (you can add more of one using ',')")
	flag.Parse()

	if *helpFlag {
		defer func() {
			recover()
		}()
		f.Println("\n-help\n\tHelp comand use to show information \n-root\n\tthe initial dir for search the delete files. \n-target\n\t The files to delete (you can add more of one using ',')\n")
		panic("")
	}

	dirToClean := *rootSearch
	dirToSwap = dirToClean

	// ! make split for get an array with the deletes targets
	targets := strings.Split(*deleteFileOrDir, ",")

	// ! celaning the targets
	for slices.Index(targets, " ") > -1 {
		indexOfEmpty := slices.Index(targets, " ")
		targets = slices.Delete(targets, indexOfEmpty, indexOfEmpty+1)
	}

	//! contains the clear targets to delete (use after the 'for' loop with this sign '#')
	cleanTargets := []string{}

	// !#
	for _, target := range targets {
		cleanTargets = append(cleanTargets, strings.Replace(target, " ", "", -1))
	}
	//! #

	recursiveNavigator(dirToClean, cleanTargets)
}

func recursiveNavigator(setterPath string, deleteTargets []string) {

	baseDirCollection := readADir(setterPath)

	for _, childDir := range baseDirCollection {

		//! name of file or dir
		elemName := childDir.Name()

		if childDir.IsDir() {

			//! path of file or dir
			parentPath := setterPath + "/" + elemName

			//! this is the perfect place to manage the files and dirs
			//* write your logic here

			// use for to loop the delete values
			for _, delTarget := range deleteTargets {
				if strings.Contains(parentPath, delTarget) {

					splitedPath := strings.Split(parentPath, "/")
					if splitedPath[len(splitedPath)-1] == delTarget {

						delErr := os.RemoveAll(parentPath)

						if delErr != nil {
							panic("An error ocurres deleting a file: " + delErr.Error())
						}
					}
				}
			}

			//* write your logic here

			// ! RECURSIVE
			recursiveNavigator(parentPath, deleteTargets)
		}

		// ! this content here only shows you the files formated
		// all this is only for my the code don't need it
		// if childDir.IsDir() { //! in case to be directory
		// 	f.Println("--------------------------------------------------")
		// 	f.Println("DirName: " + childDir.Name())
		// } else { //! in case to be a file with extension
		// 	splitedFileName := strings.Split(childDir.Name(), ".")
		// 	if len(splitedFileName) > 1 {
		// 		f.Println("--------------------------------------------------")
		// 		f.Println("FileName: " + splitedFileName[0] + "\next: " + splitedFileName[1])
		// 	}
		// }

	}
}

func readADir(path string) []fs.DirEntry {
	dir, err := os.ReadDir(path)
	h.ManageErr(err)

	return dir
}
