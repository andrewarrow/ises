package soeasy

import gc "github.com/rthornton128/goncurses"
import "time"

type SoEasyClient struct {
	s       *gc.Window
	x       int
	y       int
	buff    []byte
	line    string
	curPos  int
	recent  []RecentRoom
	curRoom int
	history []string
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
	sec.recent = recentDefaults()
	sec.curRoom = 0
	return &sec
}

func (sec *SoEasyClient) historyThread() {
	for {
		time.Sleep(time.Millisecond * 100)
		display := sec.history[0:len(sec.history)]
		limit := sec.y - 5
		if len(sec.history) > limit {
			display = sec.history[len(sec.history)-limit : len(sec.history)]
		}
		for r, oneline := range display {
			if len(oneline) > 80 {
				oneline = oneline[0:80]
			}
			sec.s.MovePrint(r, 0, oneline+"                                                                                                              ")
		}
		sec.Paint()
	}
}

func (sec *SoEasyClient) Paint() {
	r := sec.recent[sec.curRoom]
	sec.s.MovePrint(sec.y-1, 0, "                                                                              ")
	sec.s.MovePrint(sec.y-1, 0, r.name+"> "+sec.line)
	sec.s.MovePrint(sec.y-1, len(r.name)+len(sec.line)+2, "")
	sec.s.Refresh()
}

func (sec *SoEasyClient) handleReturn() bool {
	if sec.line == "quit" {
		return true
	}
	sec.history = append(sec.history, sec.line)
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

func (sec *SoEasyClient) paintCurrentRoom() {
	sec.history = roomHistoryFromCache(sec.recent[sec.curRoom].fullName)
	sec.Paint()
}

func (sec *SoEasyClient) handleNextRoom() {
	sec.curRoom++
	if sec.curRoom >= len(sec.recent) {
		sec.curRoom = 0
	}
	sec.paintCurrentRoom()
}

func (sec *SoEasyClient) handlePrevRoom() {
	sec.curRoom--
	if sec.curRoom < 0 {
		sec.curRoom = len(sec.recent) - 1
	}
	sec.paintCurrentRoom()
}

func (sec *SoEasyClient) InputLoop() {
	go sec.historyThread()
	sec.paintCurrentRoom()
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
