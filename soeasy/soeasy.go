package soeasy

import gc "github.com/rthornton128/goncurses"

type SoEasyClient struct {
	s    *gc.Window
	x    int
	y    int
	buff []byte
	line string
}

func NewSoEasyClient() *SoEasyClient {
	sec := SoEasyClient{}
	sec.s, _ = gc.Init()
	sec.y, sec.x = sec.s.MaxYX()
	gc.Echo(false)
	//gc.Raw(true)
	sec.s.Keypad(true)
	sec.buff = make([]byte, 0)
	sec.line = ""
	return &sec
}

func (sec *SoEasyClient) Paint() {
	sec.s.MovePrint(sec.y-1, 0, "> "+sec.line)
	sec.s.Refresh()
}

func (sec *SoEasyClient) handleReturn() bool {
	if sec.line == "quit" {
		return true
	}
	return false
}

func (sec *SoEasyClient) addCharToBuffer(c gc.Key) {
	sec.buff = append(sec.buff, byte(c))
	sec.line = string(sec.buff)
	//curPos++
	sec.Paint()
}

func (sec *SoEasyClient) InputLoop() {
	for {
		c := sec.s.GetChar()
		nice := gc.KeyString(c)
		if c == 10 || c == 13 {
			shouldBreak := sec.handleReturn()
			if shouldBreak == true {
				break
			}
		} else if c == 93 { // ] for next
			//handleNextRoom()
		} else if nice == "up" {
		} else if nice == "down" {
		} else if nice == "left" {
			//curPos--
		} else if nice == "right" {
			//curPos++
		} else if c == 127 { // backspace
			//handleBackspace()
		} else {
			sec.addCharToBuffer(c)
		}
	}
	gc.End()
}
