package main

import "github.com/andrewarrow/ises/room"
import "os"

func handleLookupMode() {
	os.Mkdir("id_lookup", os.ModePerm)
	os.Mkdir("reverse_lookup", os.ModePerm)
	teams := room.GetTeams()
	for _, team := range teams {
		os.Mkdir("id_lookup/"+team.Index, os.ModePerm)
		os.Mkdir("reverse_lookup/"+team.Index, os.ModePerm)
		rooms := team.Rooms()
		for _, r := range rooms {
			fstr1 := "id_lookup/" + team.Index + "/" + r["id"]
			fstr2 := "reverse_lookup/" + team.Index + "/" + r["room"]
			f, _ := os.OpenFile(fstr1, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(r["room"])
			f.Close()
			f, _ = os.OpenFile(fstr2, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(r["id"])
			f.Close()
		}
		users, _ := team.Api.GetUsers()
		for _, user := range users {
			fstr1 := "id_lookup/" + team.Index + "/" + user.ID
			fstr2 := "reverse_lookup/" + team.Index + "/" + user.Name
			f, _ := os.OpenFile(fstr1, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(user.Name)
			f.Close()
			f, _ = os.OpenFile(fstr2, os.O_CREATE|os.O_WRONLY, 0660)
			_, _ = f.WriteString(user.ID)
			f.Close()
		}
	}
}
