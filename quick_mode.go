package main

import "github.com/andrewarrow/ises/room"
import "fmt"

func handleQuickMode() {
	rooms := room.Recents()

	for _, r := range rooms {
		fmt.Println("room ", r)
	}
}
