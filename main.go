package main

import "fmt"

//import gc "github.com/rthornton128/goncurses"
//import "time"
import "os"

//import "sort"
//import "github.com/nlopes/slack"

//import "github.com/andrewarrow/ises/room"
import "github.com/andrewarrow/ises/soeasy"

//import "strings"

var (
	IsesRoot string
	history  []string
	rid      string
	realId   string
	curRoom  int
	recent   map[string]int64
)

/*
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
		for r, oneline := range display {
			if len(oneline) > 80 {
				oneline = oneline[0:80]
			}
			stdscr.MovePrint(r, 0, oneline+"                                                                                                              ")
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
					line := who + "| " + ev.Msg.Text
					wrap(line, &history)
				} else {
					who := room.IdToString(ev.Msg.User, team)
					line := name + "]" + who + "| " + ev.Msg.Text
					wrap(line, &history)
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

}

func log(str string) {
	fstr := "log.log"
	f, _ := os.OpenFile(fstr, os.O_APPEND|os.O_WRONLY, 0660)
	_, _ = f.WriteString(str + "\n")
	f.Close()
}
*/

func main() {
	IsesRoot = os.Getenv("ISES_ROOT")

	/*
		args := os.Args[1:]
		if len(args) == 0 {
			fmt.Println("./ises rid")
			return
		}
			recent = make(map[string]int64)
			recent["1_for_andrew"] = time.Now().Unix()
			recent["0_aa"] = time.Now().Unix()
			recent["2_jasoncarulli"] = time.Now().Unix()
			recent["1_general"] = time.Now().Unix()
			tokens := strings.Split(args[0], ".")
			team_str := tokens[0]
			rid = tokens[1]
			realId = room.StringToId(rid, team_str)

			var team room.Team
			teams := room.GetTeams()
			for _, t := range teams {
				go t.Rtm.ManageConnection()
				//go handleRtmInCurses(t.Rtm, t.Index)
				if team_str != t.Index {
					continue
				}
				team = t
			}

			history = make([]string, 0)
			//history = append(history, team_str)
			//history = append(history, rid)
	*/
	var client *soeasy.SoEasyClient
	client = soeasy.NewSoEasyClient()
	client.Paint()
	client.InputLoop()
	//go thready(row)

	/*
		curRoom = 0
		for {
			c := stdscr.GetChar()
			//stdscr.MovePrintf(15, 10, "|%d|", c)
			if c == 10 || c == 13 {
				team.Say(realId, line)
			} else if c == 93 {
				list := findRecents()
				curRoom++
				if curRoom >= len(list) {
					curRoom = 0
				}
				tmp := list[curRoom].name
				history = roomHistoryFromCache(tmp)
				log(tmp)
				log(fmt.Sprintf("%d", len(history)))
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
			} else {
			}
		}
	*/
}
