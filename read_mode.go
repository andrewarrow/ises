package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "sort"
import "time"
import "bufio"
import "os"
import "github.com/fatih/color"

type Cache struct {
	number   int64
	filename string
}

type ByAge []Cache

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].number < a[j].number }

func handleReadMode() {
	e_team := os.Getenv("TEAM")
	e_room := os.Getenv("ROOM")

	list := make([]Cache, 0)
	files, _ := ioutil.ReadDir(IsesRoot + "/cache/")
	for _, f := range files {
		//fmt.Println(f.Name())
		_, err := os.Stat(IsesRoot + "/cache/" + f.Name() + "/mute")
		if !os.IsNotExist(err) {
			//fmt.Println("MUTE ", f.Name())
			continue
		}

		if e_room != "" && f.Name() != fmt.Sprintf("%s_%s", e_team, e_room) {
			continue
		}

		subfiles, _ := ioutil.ReadDir(IsesRoot + "/cache/" + f.Name())
		for _, sub := range subfiles {
			tokens := strings.Split(sub.Name(), "_")
			//fmt.Println(tokens[0])
			subtokens := strings.Split(tokens[0], ".")
			number, _ := strconv.ParseInt(subtokens[0], 10, 0)
			c := Cache{}
			c.number = number
			c.filename = f.Name() + "/" + sub.Name()
			list = append(list, c)
			//t := time.Unix(number, 0)
			//fmt.Println(t)
		}
	}
	sort.Sort(ByAge(list))
	d := color.New(color.FgCyan, color.Bold)
	for _, c := range list {
		t := time.Unix(c.number, 0)
		color.Magenta(fmt.Sprintf("%s", t))
		d.Printf("   " + c.filename + "\n")

		f, _ := os.Open(IsesRoot + "/cache/" + c.filename)
		//fmt.Println("wow ", err)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println("      " + scanner.Text())
		}
		f.Close()
	}
}
