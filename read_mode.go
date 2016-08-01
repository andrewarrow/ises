package main

import "fmt"
import "io/ioutil"

func handleReadMode() {
	files, _ := ioutil.ReadDir("cache/")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
