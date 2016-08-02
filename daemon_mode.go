package main

import "os"
import "fmt"
import "github.com/nlopes/slack"
import "github.com/andrewarrow/ises/room"
import "time"

func handleDaemonMode() {
	err := os.Mkdir("ui", os.ModePerm)
	if err != nil {
		fmt.Println("mkdir ", err)
	}

	teams := room.GetTeams()
	for _, team := range teams {
		go team.Rtm.ManageConnection()
		go handleRtm(team.Rtm, team.Index)
	}

	for {
		time.Sleep(time.Second)
	}
}

func handleRtm(rtm *slack.RTM, team string) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello
			case *slack.MessageEvent:
				name := room.IdToString(ev.Msg.Channel, team)
				h := make(map[string]string)
				h["text"] = ev.Msg.Text
				h["time"] = ev.Msg.Timestamp
				h["who"] = ev.Msg.User
				filename := fmt.Sprintf("%s_%s", team, name)
				room.WriteMessageToDisk(filename, h)
				fmt.Printf("Message: %s %s\n", name, team)
			case *slack.PresenceChangeEvent:
				name := room.IdToString(ev.User, team)
				fmt.Printf("Presence Change: %s %s\n", name, team)
			}
		}
	}

}
