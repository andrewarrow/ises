package soeasy

import "os"
import "fmt"
import "github.com/andrewarrow/ises/room"

func SoEasySetup() {
	err := os.Mkdir("cache", os.ModePerm)
	if err != nil {
		fmt.Println("Could not create dir: ", err.Error())
		return
	}
	teams := room.GetTeams()
	if len(teams) == 0 {
		fmt.Println("Set your SLACK_TEAMS=1 and SLACK_TOKEN_0 env vars")
		return
	}
	for _, team := range teams {
		fmt.Println("Caching users from team ", team.Index)
		cache_users(team)
		fmt.Println("Caching rooms from team ", team.Index)
		cache_rooms(team)
	}

}

func cache_rooms(team room.Team) {
	rooms := team.Rooms()
	for _, r := range rooms {
		fstr1 := "cache/id_lookup/" + team.Index + "/" + r["id"]
		fstr2 := "cache/reverse_lookup/" + team.Index + "/" + r["room"]
		f, _ := os.OpenFile(fstr1, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(r["room"])
		f.Close()
		f, _ = os.OpenFile(fstr2, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(r["id"])
		f.Close()
	}
}

func cache_users(team room.Team) {
	users, _ := team.Api.GetUsers()
	for _, user := range users {
		dir1 := "cache/id_lookup/" + team.Index + "/"
		dir2 := "cache/reverse_lookup/" + team.Index + "/"

		_ = os.MkdirAll(dir1, os.ModePerm)
		_ = os.MkdirAll(dir2, os.ModePerm)

		fstr1 := dir1 + user.ID
		fstr2 := dir2 + user.Name
		f, _ := os.OpenFile(fstr1, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(user.Name)
		f.Close()
		f, _ = os.OpenFile(fstr2, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(user.ID)
		f.Close()
	}
}
