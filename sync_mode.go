package main

import "github.com/andrewarrow/ises/room"
import "fmt"
import "os"
import "bufio"
import "strings"

func handleSyncMode() {
	teams := room.GetTeams()
	for _, team := range teams {
		rooms := team.Rooms()
		for _, r := range rooms {
			if r["room"] == "slackbot" {
				continue
			}
			filename := fmt.Sprintf("%s_%s", team.Index, r["room"])
			fmt.Println(filename)
			handleFile(filename, team, r)
		}
	}
}

func findLatest(filename string) string {
	f, err := os.Open("cache/" + filename)
	fmt.Println("findLatest ", err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	last := ""
	for scanner.Scan() {
		last = scanner.Text()
	}
	if last == "" {
		return ""
	}
	tokens := strings.Split(last, "|")
	fmt.Println(tokens[0])
	return tokens[0]
}

func handleFile(filename string, team room.Team, room map[string]string) {
	latest := "" //findLatest(filename)
	history := team.History(room["id"], room["thing"], latest)
	if len(history) == 0 {
		return
	}

	err := os.Mkdir("cache/"+filename, os.ModePerm)
	fmt.Println("mkdir ", err)

	i := len(history) - 1
	for {
		h := history[i]

		fstr := "cache/" + filename + "/" + h["time"] + "_" + h["who"]
		_, err := os.Stat(fstr)
		if os.IsNotExist(err) {
			f, err := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0600)
			fmt.Println("open file ", err)
			defer f.Close()

			_, err = f.WriteString(h["text"])
			fmt.Println("f.WriteString ", err)
		}

		i--
		if i < 0 {
			break
		}
	}
}
