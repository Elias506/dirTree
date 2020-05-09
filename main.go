package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	path := os.Args[1]
	tabs := ""
	err := dirTree(out, path, tabs)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, tabs string) error {
	fileObj, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", path, err.Error())
	}
	defer fileObj.Close()
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatalf("Could not open dir %s: %s\n\n", fileName, err.Error())
	}

	for _, file:= range files {
		var tabs2 string
		if files[len(files)-1] == file {
			tabs2 = tabs + "└"
		} else {
			tabs2 = tabs + "├"
		}
		if file.IsDir() {
			newDir := filepath.Join(path, file.Name())
			fmt.Fprintf(out, "%s───%s\n", tabs2, file.Name())
			var tabs1 string
			if file == files[len(files)-1]{
				tabs1 = tabs + "    "
			} else {
				tabs1 = tabs + "│   "
			}
			dirTree(out, newDir, tabs1)
		} else {
			if file.Size() > 0 {
				fmt.Fprintf(out, "%s───%s (%vb)\n", tabs2, file.Name(), file.Size())
			} else {
				fmt.Fprintf(out, "%s───%s (empty)\n", tabs2, file.Name())
			}
		}
	}
	return nil
}

