package main

import "fmt"
import gc "github.com/rthornton128/goncurses"
import "time"

var (
	IsesRoot string
	stdscr   *gc.Window
	history  []string
)

func thready() {
	for {
		time.Sleep(time.Second)
		for row, h := range history {
			stdscr.MovePrint(row, 0, h)
		}
		stdscr.Refresh()
	}
}

func main() {
	history = make([]string, 0)
	stdscr, _ = gc.Init()
	go thready()
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
			//stdscr.MovePrintf(10, 10, "%s", line)
			history = append(history, line)
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
