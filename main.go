package main

import "fmt"
import gc "github.com/rthornton128/goncurses"

var (
	IsesRoot string
)

func main() {
	stdscr, _ := gc.Init()
	row, _ := stdscr.MaxYX()
	defer gc.End()

	stdscr.MovePrint(row-1, 0, "> ")
	stdscr.Refresh()
	buff := make([]byte, 0)
	line := ""
	for {
		c := stdscr.GetChar()
		if c == 10 || c == 13 {
			stdscr.MovePrintf(10, 10, "%s", line)
			stdscr.MovePrint(row-1, 0, "> ")
			stdscr.Refresh()
			if line == "quit" {
				fmt.Println("")
				break
			}
		}
		buff = append(buff, byte(c))
		line = string(buff)
	}
}
