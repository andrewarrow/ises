package soeasy

import gc "github.com/rthornton128/goncurses"
import "time"
import "github.com/andrewarrow/ises/room"
import "sync"

//import "fmt"

var EasyMutex = &sync.Mutex{}

type SoEasyClient struct {
	s            *gc.Window
	x            int
	y            int
	buff         []byte
	line         string
	curPos       int
	recent       []RecentRoom
	curRoomIndex int
	curRoom      RecentRoom
	history      []string
	teams        []room.Team
	team         room.Team
}

func NewSoEasyClient() *SoEasyClient {
	sec := SoEasyClient{}
	sec.s, _ = gc.Init()
	sec.y, sec.x = sec.s.MaxYX()
	gc.Echo(false)
	//gc.Raw(true)
	sec.s.Keypad(true)
	sec.curPos = 0
	sec.buff = make([]byte, 0)
	sec.history = make([]string, 0)
	sec.line = ""
	sec.recent = make([]RecentRoom, 0)
	recentDefaults(&sec)
	sec.curRoomIndex = 0
	sec.curRoom = sec.recent[0]
	return &sec
}

func (sec *SoEasyClient) addToRecentOrUpdateTs(r RecentRoom) {
	for _, tmp := range sec.recent {
		if tmp.fullName == r.fullName {
			tmp.ts = r.ts
			return
		}
	}
	// TODO sort by ts
	sec.recent = append(sec.recent, r)
}

func (sec *SoEasyClient) historyThread() {
	for {
		time.Sleep(time.Millisecond * 100)
		display := sec.history[0:len(sec.history)]
		limit := sec.y - 5
		if len(sec.history) > limit {
			display = sec.history[len(sec.history)-limit : len(sec.history)]
		}
		EasyMutex.Lock()
		for r, oneline := range display {
			if len(oneline) > 80 {
				oneline = oneline[0:80]
			}
			sec.s.MovePrint(r, 0, oneline+"                                                                                                              ")
		}
		EasyMutex.Unlock()
		sec.Paint()
	}
}

func (sec *SoEasyClient) Paint() {
	r := sec.curRoom
	EasyMutex.Lock()
	sec.s.MovePrint(sec.y-1, 0, "                                                                              ")
	sec.s.MovePrint(sec.y-1, 0, r.team+"_"+r.name+"> "+sec.line)
	sec.s.MovePrint(sec.y-1, 2+len(r.name)+len(sec.line)+2, "")
	sec.s.Refresh()
	EasyMutex.Unlock()
}

func (sec *SoEasyClient) handleReturn() bool {
	if sec.line == "quit" {
		return true
	}
	sec.team.Say(sec.curRoom.realId, sec.line)
	sec.buff = make([]byte, 0)
	sec.line = ""
	sec.curPos = 0
	sec.Paint()
	return false
}

func (sec *SoEasyClient) addCharToBuffer(c gc.Key) {
	sec.buff = append(sec.buff, byte(c))
	sec.line = string(sec.buff)
	//curPos++
	sec.Paint()
}

func (sec *SoEasyClient) handleBackspace() {
	if len(sec.buff) > 0 {
		sec.buff = sec.buff[0 : len(sec.buff)-1]
		sec.curPos--
		sec.line = string(sec.buff)
		sec.Paint()
	}
}

func (sec *SoEasyClient) roomChange() {
	sec.curRoom = sec.recent[sec.curRoomIndex]
	sec.history = roomHistoryFromCache(sec.curRoom.fullName)
	sec.team = sec.teams[sec.curRoom.teamIndex]
	y := 0
	EasyMutex.Lock()
	for {
		sec.s.MovePrint(y, 0, "                                                                                                              ")
		y++
		if y > sec.y-5 {
			break
		}
	}
	EasyMutex.Unlock()
	sec.Paint()
}

func (sec *SoEasyClient) handleNextRoom() {
	sec.curRoomIndex++
	if sec.curRoomIndex >= len(sec.recent) {
		sec.curRoomIndex = 0
	}
	sec.roomChange()
}

func (sec *SoEasyClient) handlePrevRoom() {
	sec.curRoomIndex--
	if sec.curRoomIndex < 0 {
		sec.curRoomIndex = len(sec.recent) - 1
	}
	sec.roomChange()
}

func (sec *SoEasyClient) setupWebsocket() {
	sec.teams = room.GetTeams()
	for _, t := range sec.teams {
		go t.Rtm.ManageConnection()
		go sec.handleRtmInCurses(t.Rtm, t.Index)
	}
}

func (sec *SoEasyClient) InputLoop() {
	go sec.historyThread()
	sec.setupWebsocket()
	go sec.lookForMissingMessages()
	sec.roomChange()
	for {
		c := sec.s.GetChar()
		nice := gc.KeyString(c)
		if c == 10 || c == 13 {
			shouldBreak := sec.handleReturn()
			if shouldBreak == true {
				break
			}
		} else if c == 91 { // [ for prev
			sec.handlePrevRoom()
		} else if c == 93 { // ] for next
			sec.handleNextRoom()
		} else if nice == "up" {
		} else if nice == "down" {
		} else if nice == "left" {
		} else if nice == "right" {
		} else if c == 127 { // backspace
			sec.handleBackspace()
		} else {
			sec.addCharToBuffer(c)
		}
	}
	gc.End()
}
