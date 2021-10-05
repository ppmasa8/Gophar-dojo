package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func convert(fpath string) string {
	str := fpath
	rep := regexp.MustCompile(`png$`)
	str = rep.ReplaceAllString(str, "jpg")
	if err := os.Rename(fpath, str); err != nil {
		log.Fatal(err)
	}
	return str
}

func main() {
	for _, f := range GetFiles("./dojo1/kadai1/sample") {
		fmt.Println("Convert", f, "to", convert(f))
	}
}

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
