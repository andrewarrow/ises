package main

import "fmt"
import gc "github.com/rthornton128/goncurses"
import "time"
import "os"
import "github.com/nlopes/slack"
import "github.com/andrewarrow/ises/room"
import "strings"

var (
	IsesRoot string
	stdscr   *gc.Window
	history  []string
	rid      string
	realId   string
	line     string
	curPos   int
)

func thready(maxrows int) {
	for {
		time.Sleep(time.Millisecond * 100)
		display := history[0:len(history)]
		limit := maxrows - 5
		if len(history) > limit {
			display = history[len(history)-limit : len(history)]
		}
		for r, h := range display {
			stdscr.MovePrint(r, 0, h)
		}
		stdscr.MovePrint(maxrows-1, curPos+2, "")
		stdscr.Refresh()
	}
}

func handleRtmInCurses(rtm *slack.RTM, team string) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				name := room.IdToString(ev.Msg.Channel, team)
				h := make(map[string]string)
				h["text"] = ev.Msg.Text
				h["time"] = ev.Msg.Timestamp
				h["who"] = ev.Msg.User
				filename := fmt.Sprintf("%s_%s", team, name)
				room.WriteMessageToDisk(filename, h)
				if name == rid {
					who := room.IdToString(ev.Msg.User, team)
					history = append(history, who+"| "+ev.Msg.Text)
				}
			}
		}
	}

}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("./ises rid")
		return
	}
	tokens := strings.Split(args[0], ".")
	team_str := tokens[0]
	rid = tokens[1]
	realId = room.StringToId(rid, team_str)

	var team room.Team
	teams := room.GetTeams()
	for _, t := range teams {
		if team_str != t.Index {
			continue
		}
		team = t
		go team.Rtm.ManageConnection()
		go handleRtmInCurses(team.Rtm, team.Index)
	}

	history = make([]string, 0)
	//history = append(history, team_str)
	//history = append(history, rid)
	stdscr, _ = gc.Init()
	row, _ := stdscr.MaxYX()
	go thready(row)
	gc.Echo(false)
	//gc.Raw(true)
	stdscr.Keypad(true)

	stdscr.MovePrint(row-1, 0, "> ")
	stdscr.Refresh()
	buff := make([]byte, 0)
	line = ""
	curPos = 0
	for {
		c := stdscr.GetChar()
		//stdscr.MovePrintf(15, 10, "|%d|", c)
		if c == 10 || c == 13 {
			//stdscr.MovePrintf(10, 10, "%s", line)
			team.Say(realId, line)
			buff = make([]byte, 0)
			curPos = 0
			stdscr.MovePrint(row-1, 0, "                                                                   ")
			stdscr.MovePrint(row-1, 0, "> ")
			stdscr.Refresh()
			if line == "quit" {
				gc.End()
				fmt.Println("")
				break
			}
		} else {
			nice := gc.KeyString(c)
			if nice == "up" {
			} else if nice == "down" {
			} else if nice == "left" {
				curPos--
			} else if nice == "right" {
				curPos++
			} else if c == 127 {
				if len(buff) > 0 {
					buff = buff[0 : len(buff)-1]
					curPos--
					line = string(buff)
					stdscr.MovePrint(row-1, 0, "> "+line+" ")
					stdscr.MovePrint(row-1, len(line)+2, "")
				}
			} else {
				buff = append(buff, byte(c))
				line = string(buff)
				curPos++
				stdscr.MovePrint(row-1, 0, "> "+line)
			}
			stdscr.Refresh()
		}
	}
}
