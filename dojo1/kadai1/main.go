package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func convert() {

}

func main() {
	// Readfiles
	files, err := ioutil.ReadDir("./dojo1/kadai1/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
