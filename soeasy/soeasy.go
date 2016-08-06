package soeasy

import gc "github.com/rthornton128/goncurses"

type SoEasyClient struct {
	stdscr *gc.Window
	x      int
	y      int
}

func NewSoEasyClient() *SoEasyClient {
	sec := SoEasyClient{}
	sec.stdscr, _ = gc.Init()
	sec.y, sec.x = sec.stdscr.MaxYX()
	gc.Echo(false)
	//gc.Raw(true)
	sec.stdscr.Keypad(true)
	return &sec
}

func (sec *SoEasyClient) Paint() {
}
