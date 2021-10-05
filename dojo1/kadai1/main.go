package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func convert() {

}

func main() {
	// Read files and dirs.
	for _, f := range GetFiles("./dojo1/kadai1/sample") {
		fmt.Println(f)
	}
}

// TODO split get files func and recursive dir
func GetFiles(dir string) []string {
	var result []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			result = append(result, GetFiles(fpath)...)
			continue
		}
		result = append(result, fpath)
	}
	return result
}
