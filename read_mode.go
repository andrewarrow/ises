package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "time"

func handleReadMode() {
	files, _ := ioutil.ReadDir("cache/")
	for _, f := range files {
		fmt.Println(f.Name())
		subfiles, _ := ioutil.ReadDir("cache/" + f.Name())
		for _, sub := range subfiles {
			tokens := strings.Split(sub.Name(), "_")
			fmt.Println(tokens[0])
			subtokens := strings.Split(tokens[0], ".")
			number, _ := strconv.ParseInt(subtokens[0], 10, 0)
			t := time.Unix(int64(number), 0)
			fmt.Println(t)
		}
	}
}
