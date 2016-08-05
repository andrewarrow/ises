package main

import "fmt"
import gc "github.com/rthornton128/goncurses"

var (
	IsesRoot string
)

func main() {
	stdscr, _ := gc.Init()
	gc.Echo(false)
	//gc.Raw(true)
	stdscr.Keypad(true)
	row, _ := stdscr.MaxYX()

	stdscr.MovePrint(row-1, 0, "> ")
	stdscr.Refresh()
	buff := make([]byte, 0)
	line := ""
	for {
		c := stdscr.GetChar()
		//stdscr.MovePrintf(15, 10, "|%d|", c)
		if c == 10 || c == 13 {
			stdscr.MovePrintf(10, 10, "%s", line)
			buff = make([]byte, 0)
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
			} else if nice == "right" {
			} else if c == 127 {
				if len(buff) > 0 {
					buff = buff[0 : len(buff)-1]
					line = string(buff)
					stdscr.MovePrint(row-1, 0, "> "+line+" ")
					stdscr.MovePrint(row-1, len(line)+2, "")
				}
			} else {
				buff = append(buff, byte(c))
				line = string(buff)
				stdscr.MovePrint(row-1, 0, "> "+line)
			}
			stdscr.Refresh()
		}
	}
}
