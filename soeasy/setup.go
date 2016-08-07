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

	}

}
