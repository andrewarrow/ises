package main

import "github.com/andrewarrow/ises/room"
import "fmt"

func handleSyncMode() {
  teams := room.GetTeams()
  for _, team := range teams {
    rooms := team.Rooms()
    for _, r := range rooms {
      filename := fmt.Sprintf("%s_%s", team.Index, r["room"])
      fmt.Println(filename)
      handleFile(filename, team, r)
    }
  }
}

func handleFile(filename string, team room.Team, room map[string]string) {
  history := team.History(room["id"])
  fmt.Println("hi", history)
}
