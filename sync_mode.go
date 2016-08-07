package main

import "github.com/andrewarrow/ises/room"
import "fmt"
import "os"
import "bufio"

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
			lookForSay(filename, team, r)
			handleFile(filename, team, r)
		}
	}
}

func lookForSay(filename string, team room.Team, r map[string]string) {
	sayfile := "cache/" + filename + "/say"
	_, err := os.Stat(sayfile)
	if os.IsNotExist(err) {
		return
	}
	f, _ := os.Open(sayfile)
	//fmt.Println("wow ", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		team.Say(r["id"], scanner.Text())
	}
	f.Close()
	os.Remove(sayfile)
}

func handleFile(filename string, team room.Team, r map[string]string) {
	_ = os.Mkdir("cache/"+filename, os.ModePerm)
	//fmt.Println("mkdir ", err)
	_, err := os.Stat("cache/" + filename + "/mute")
	if !os.IsNotExist(err) {
		fmt.Println("MUTE ", filename)
		return
	}

	history := team.History(r["id"], r["thing"])
	if len(history) == 0 {
		return
	}

	i := len(history) - 1
	for {
		h := history[i]

		room.WriteMessageToDisk(filename, h)

		i--
		if i < 0 {
			break
		}
	}
}
