package soeasy

import "io/ioutil"
import "strings"
import "strconv"
import "sort"
import "github.com/andrewarrow/ises/room"
import "fmt"

//import "time"
import "bufio"
import "os"

type Cache struct {
	number   int64
	filename string
	fullDate string
}

type ByAge []Cache

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].number < a[j].number }

type ByAgeRev []Cache

func (a ByAgeRev) Len() int           { return len(a) }
func (a ByAgeRev) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAgeRev) Less(i, j int) bool { return a[i].number > a[j].number }

func log(str string) {
	fstr := "log.log"
	f, _ := os.OpenFile(fstr, os.O_APPEND|os.O_WRONLY, 0660)
	_, _ = f.WriteString(str + "\n")
	f.Close()
}

func (sec *SoEasyClient) lookForMissingMessages() {
	list := allRoomsStep1()
	already := make(map[string]string)
	teams := room.GetTeams(false)
	for _, c := range list {
		tokens := strings.Split(c.filename, "/")
		room_name := tokens[0]
		_, err := os.Stat("cache/messages/" + room_name + "/mute")
		if !os.IsNotExist(err) {
			continue
		}
		if already[room_name] == "" {
			teamStr := room_name[0:1]
			teamIndex, _ := strconv.Atoi(teamStr)
			realId := room.StringToId(room_name[2:], teamStr)
			missing := teams[teamIndex].History(realId, realId[0:1], c.fullDate)
			for _, h := range missing {
				room.WriteMessageToDisk(room_name, h)
				log(room_name)
				log(fmt.Sprintf("%v", h))
			}
			already[room_name] = "1"
		}
	}
}

func allRoomsStep1() []Cache {
	list := make([]Cache, 0)
	subfiles, _ := ioutil.ReadDir("cache/messages/")
	for _, sub := range subfiles {
		for _, a := range roomHistoryStep1(sub.Name()) {
			list = append(list, a)
		}
	}
	sort.Sort(ByAgeRev(list))
	return list
}

func computeLatestRooms() []string {

	res := make([]string, 0)
	list := allRoomsStep1()

	already := make(map[string]string)
	for _, c := range list {
		tokens := strings.Split(c.filename, "/")
		room_name := tokens[0]
		_, err := os.Stat("cache/messages/" + room_name + "/mute")
		if !os.IsNotExist(err) {
			continue
		}

		already[room_name] = "1"

		if len(already) > 10 {
			break
		}
	}
	for k, _ := range already {
		res = append(res, k)
	}

	return res
}

func roomHistoryStep1(room_file string) []Cache {

	list := make([]Cache, 0)
	subfiles, _ := ioutil.ReadDir("cache/messages/" + room_file)
	for _, sub := range subfiles {
		if sub.Name() == "mute" {
			continue
		}
		tokens := strings.Split(sub.Name(), "_")

		subtokens := strings.Split(tokens[0], ".")
		number, _ := strconv.ParseInt(subtokens[0], 10, 0)
		c := Cache{}
		c.number = number
		c.fullDate = tokens[0]
		c.filename = room_file + "/" + sub.Name()
		list = append(list, c)
	}
	return list
}

func roomHistoryFromCache(room_file string) []string {

	list := roomHistoryStep1(room_file)

	sort.Sort(ByAge(list))
	history := make([]string, 0)
	for _, c := range list {
		//t := time.Unix(c.number, 0)
		f, _ := os.Open("cache/messages/" + c.filename)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			wrap(line, &history)
		}
		f.Close()
	}
	return history
}

func wrap(line string, his *[]string) {
	for {
		if len(line) > 80 {
			*his = append(*his, line[0:80])
			line = line[80:len(line)]
		} else {
			break
		}
	}
	*his = append(*his, line)
}
