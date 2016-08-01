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
				fmt.Printf("Message: %v %s\n", ev, team)
			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v %s\n", ev, team)
			}
		}
	}

}