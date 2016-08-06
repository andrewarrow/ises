package main

//import "fmt"

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

	/*
		args := os.Args[1:]
		if len(args) == 0 {
			fmt.Println("./ises rid")
			return
		}
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
