package main

import (
	f "fmt"
	"os"

	h "github.com/ManuelsSaNt/goCleaner/helpers"
)

func main() {

	f.Println("mamalon")

	readDir, err := os.ReadDir("../")

	h.ManageErr(err)

	for _, dirOrFile := range readDir {
		f.Println(dirOrFile)
	}

}
