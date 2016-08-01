package main

import "github.com/andrewarrow/ises/room"
import "os"

func handleLookupMode() {
	os.Mkdir("id_lookup", os.ModePerm)
	teams := room.GetTeams()
	for _, team := range teams {
		os.Mkdir("id_lookup/"+team.Index, os.ModePerm)
		rooms := team.Rooms()
		for _, r := range rooms {
			fstr := "id_lookup/" + team.Index + "/" + r["id"]
			f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(r["room"])
			f.Close()
		}
		users, _ := team.Api.GetUsers()
		for _, user := range users {
			fstr := "id_lookup/" + team.Index + "/" + user.ID
			f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(user.Name)
			f.Close()
		}
	}
}
