package main

import "io/ioutil"
import "strings"
import "strconv"
import "sort"

//import "time"
import "bufio"
import "os"

type Cache struct {
	number   int64
	filename string
}

type ByAge []Cache

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].number < a[j].number }

func roomHistoryFromCache(room_file string) []string {

	list := make([]Cache, 0)
	subfiles, _ := ioutil.ReadDir(IsesRoot + "/cache/" + room_file)
	for _, sub := range subfiles {
		tokens := strings.Split(sub.Name(), "_")
		subtokens := strings.Split(tokens[0], ".")
		number, _ := strconv.ParseInt(subtokens[0], 10, 0)
		c := Cache{}
		c.number = number
		c.filename = room_file + "/" + sub.Name()
		list = append(list, c)
	}

	sort.Sort(ByAge(list))
	history := make([]string, 0)
	for _, c := range list {
		//t := time.Unix(c.number, 0)
		f, _ := os.Open(IsesRoot + "/cache/" + c.filename)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			//fmt.Println("      " + scanner.Text())
			history = append(history, scanner.Text())
		}
		f.Close()
	}
	return history
}
