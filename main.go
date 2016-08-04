package main

import "fmt"
import gc "github.com/rthornton128/goncurses"

var (
	IsesRoot string
)

func main() {
	stdscr, _ := gc.Init()
	row, _ := stdscr.MaxYX()

	stdscr.MovePrint(row-1, 0, "> ")
	stdscr.Refresh()
	buff := make([]byte, 0)
	line := ""
	for {
		c := stdscr.GetChar()
		if c == 10 || c == 13 {
			buff = make([]byte, 0)
			stdscr.MovePrintf(10, 10, "|%s|%d", line, len(line))
			stdscr.MovePrint(row-1, 0, "> ")
			stdscr.Refresh()
			if line == "quit" {
				gc.End()
				fmt.Println("")
				break
			}
		} else {
			buff = append(buff, byte(c))
			line = string(buff)
		}
	}
}
