package soeasy

import "github.com/andrewarrow/ises/room"
import "github.com/nlopes/slack"
import "time"
import "fmt"

func (sec *SoEasyClient) handleRtmInCurses(rtm *slack.RTM, team string) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				name := room.IdToString(ev.Msg.Channel, team)
				r := NewRecentRoom(team + "_" + name)
				r.ts = time.Now().Unix()
				r.realId = ev.Msg.Channel
				sec.addToRecentOrUpdateTs(r)

				h := make(map[string]string)
				h["text"] = ev.Msg.Text
				h["time"] = ev.Msg.Timestamp
				h["who"] = ev.Msg.User
				filename := fmt.Sprintf("%s_%s", team, name)
				room.WriteMessageToDisk(filename, h)

				if name == sec.curRoom.name {
					who := room.IdToString(ev.Msg.User, team)
					line := who + "| " + ev.Msg.Text
					wrap(line, &sec.history)
				} else {
					who := room.IdToString(ev.Msg.User, team)
					line := name + "]" + who + "| " + ev.Msg.Text
					wrap(line, &sec.history)
				}
			}
		}
	}

}
