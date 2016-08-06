package soeasy

import gc "github.com/rthornton128/goncurses"

type SoEasyClient struct {
	s *gc.Window
	x int
	y int
}

func NewSoEasyClient() *SoEasyClient {
	sec := SoEasyClient{}
	sec.s, _ = gc.Init()
	sec.y, sec.x = sec.s.MaxYX()
	gc.Echo(false)
	//gc.Raw(true)
	sec.s.Keypad(true)
	return &sec
}

func (sec *SoEasyClient) Paint() {
	sec.s.MovePrint(sec.y-1, 0, "> ")
	sec.s.Refresh()
}

func (sec *SoEasyClient) InputLoop() {
	for {
		c := sec.s.GetChar()
		if c == 10 || c == 13 {
			break
		}
	}
	gc.End()
}
