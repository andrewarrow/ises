package main

import "github.com/andrewarrow/ises/room"
import "fmt"
import "os"
import "bufio"

func handleSyncMode() {
	teams := room.GetTeams()
	for _, team := range teams {
		rooms := team.Rooms()
		for _, r := range rooms {
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return ""
}

func handleFile(filename string, team room.Team, room map[string]string) {
	latest := findLatest(filename)
	history := team.History(room["id"], room["thing"], latest)
	if len(history) == 0 {
		return
	}

	f, err := os.OpenFile("cache/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	fmt.Println("open file ", err)
	defer f.Close()

	i := len(history) - 1
	for {
		h := history[i]
		str := fmt.Sprintf("%s|%s|%s\n", h["time"], h["who"], h["text"])
		_, err = f.WriteString(str)
		fmt.Println("f.WriteString ", err)
		i--
		if i < 0 {
			break
		}
	}
}
