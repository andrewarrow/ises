package main

import "fmt"
import gc "github.com/rthornton128/goncurses"
import "time"
import "os"
import "sort"
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
	curRoom  int
	recent   map[string]int64
)

type RecentRoom struct {
	ts   int64
	name string
}

type ByRoomAge []RecentRoom

func (a ByRoomAge) Len() int           { return len(a) }
func (a ByRoomAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRoomAge) Less(i, j int) bool { return a[i].ts < a[j].ts }

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
		stdscr.MovePrint(maxrows-1, curPos+2+len(rid), "")
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
				recent[team+"_"+name] = time.Now().Unix()
				h := make(map[string]string)
				h["text"] = ev.Msg.Text
				h["time"] = ev.Msg.Timestamp
				h["who"] = ev.Msg.User
				filename := fmt.Sprintf("%s_%s", team, name)
				room.WriteMessageToDisk(filename, h)
				if name == rid {
					who := room.IdToString(ev.Msg.User, team)
					history = append(history, who+"| "+ev.Msg.Text)
				} else {
					who := room.IdToString(ev.Msg.User, team)
					history = append(history, name+"]"+who+"| "+ev.Msg.Text)
				}
			}
		}
	}

}

func findRecents() []RecentRoom {
	list := make([]RecentRoom, 0)
	for k, v := range recent {
		rr := RecentRoom{}
		rr.ts = v
		rr.name = k
		list = append(list, rr)
	}
	sort.Sort(ByRoomAge(list))
	return list

	/*
		fstr := "log.log"
		f, _ := os.OpenFile(fstr, os.O_APPEND|os.O_WRONLY, 0660)
		for _, item := range list {
			_, _ = f.WriteString(item.name + "\n")
		}
		f.Close()
	*/
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("./ises rid")
		return
	}
	recent = make(map[string]int64)
	recent["1_for_andrew"] = time.Now().Unix()
	recent["0_aa"] = time.Now().Unix()
	recent["2_jasoncarulli"] = time.Now().Unix()
	tokens := strings.Split(args[0], ".")
	team_str := tokens[0]
	rid = tokens[1]
	realId = room.StringToId(rid, team_str)

	var team room.Team
	teams := room.GetTeams()
	for _, t := range teams {
		go t.Rtm.ManageConnection()
		go handleRtmInCurses(t.Rtm, t.Index)
		if team_str != t.Index {
			continue
		}
		team = t
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

	stdscr.MovePrint(row-1, 0, rid+"> ")
	stdscr.Refresh()
	buff := make([]byte, 0)
	line = ""
	curPos = 0
	curRoom = 0
	for {
		c := stdscr.GetChar()
		//stdscr.MovePrintf(15, 10, "|%d|", c)
		if c == 10 || c == 13 {
			if line == "quit" {
				gc.End()
				fmt.Println("")
				break
			}
			team.Say(realId, line)
			buff = make([]byte, 0)
			curPos = 0
			stdscr.MovePrint(row-1, 0, "                                                                   ")
			stdscr.MovePrint(row-1, 0, rid+"> ")
			stdscr.Refresh()
		} else if c == 93 {
			list := findRecents()
			curRoom++
			if curRoom >= len(list) {
				curRoom = 0
			}
			tmp := list[curRoom].name
			rid = tmp[2:len(tmp)]
			newTeam := tmp[0:1]
			realId = room.StringToId(rid, newTeam)
			teams := room.GetTeams()
			for _, t := range teams {
				if newTeam != t.Index {
					continue
				}
				team = t
			}
			curPos = 0
			stdscr.MovePrint(row-1, 0, "                                                                   ")
			stdscr.MovePrint(row-1, 0, rid+"> ")
			stdscr.Refresh()
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
					stdscr.MovePrint(row-1, 0, rid+"> "+line+" ")
					stdscr.MovePrint(row-1, len(line)+2+len(rid), "")
				}
			} else {
				buff = append(buff, byte(c))
				line = string(buff)
				curPos++
				stdscr.MovePrint(row-1, 0, rid+"> "+line)
			}
			stdscr.Refresh()
		}
	}
}
