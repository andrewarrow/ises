package main

import "github.com/andrewarrow/ises/room"
import "fmt"
import "os"

func handleSyncMode() {
	err := os.Mkdir("cache", os.ModePerm)
	fmt.Println("mkdir ", err)

	teams := room.GetTeams()
	for _, team := range teams {
		rooms := team.Rooms()
		for _, r := range rooms {
			if r["room"] == "slackbot" {
				continue
			}
			filename := fmt.Sprintf("%s_%s", team.Index, r["room"])
			//fmt.Println(filename)
			handleFile(filename, team, r)
		}
	}
}

func handleFile(filename string, team room.Team, room map[string]string) {
	_ = os.Mkdir("cache/"+filename, os.ModePerm)
	//fmt.Println("mkdir ", err)
	_, err := os.Stat("cache/" + filename + "/mute")
	if !os.IsNotExist(err) {
		fmt.Println("MUTE ", filename)
		return
	}

	history := team.History(room["id"], room["thing"], "")
	if len(history) == 0 {
		return
	}

	i := len(history) - 1
	for {
		h := history[i]

		fstr := "cache/" + filename + "/" + h["time"] + "_" + h["who"]
		_, err := os.Stat(fstr)
		if os.IsNotExist(err) {
			f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0600)
			//fmt.Println("open file ", err)
			_, _ = f.WriteString(h["text"])
			//fmt.Println("f.WriteString ", err)
			f.Close()
		}

		i--
		if i < 0 {
			break
		}
	}
}
