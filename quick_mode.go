package main

import "github.com/andrewarrow/ises/room"
import "fmt"

func handleQuickMode() {
  teams := room.GetTeams()
  for _, team := range teams {
    recents := team.Recents()
    for _, r := range recents {
      s := fmt.Sprintf("%s %s(%s)", team.Index, r["room"], r["count"])
      fmt.Println(s)
    }
  }
}
