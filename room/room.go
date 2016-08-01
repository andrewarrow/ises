package room

import "github.com/nlopes/slack"
import "strconv"
import "os"
import "fmt"

type Team struct {
   Api *slack.Client
   Index string
}

func (t Team)Recents() []map[string]string {
  list := make([]map[string]string, 0)
  if t.Index == "0" {
	  m := make(map[string]string)
	  m["room"] = "ac_biz"
	  m["count"] = "4"
	  list = append(list, m)
	  m = make(map[string]string)
	  m["room"] = "marketing"
	  m["count"] = "3"
	  list = append(list, m)
  }
  if t.Index == "1" {
	  m := make(map[string]string)
	  m["room"] = "payouts"
	  m["count"] = "1"
	  list = append(list, m)
  }

  return list
}

func (t Team)History(id string) []map[string]string {
  list := make([]map[string]string, 0)
  if id == "123" {
    m := make(map[string]string)
    m["text"] = "hello"
    m["time"] = "08:03"
    m["who"] = "bob"
    list = append(list, m)
  }
  return list
}

func (t Team)Rooms() []map[string]string {
  list := make([]map[string]string, 0)
  if t.Index == "0" {
	  m := make(map[string]string)
	  m["room"] = "ac_biz"
	  m["id"] = "123"
	  list = append(list, m)
	  m = make(map[string]string)
	  m["room"] = "marketing"
	  m["id"] = "124"
	  list = append(list, m)
	  m = make(map[string]string)
	  m["room"] = "tanguy"
	  m["id"] = "125"
	  list = append(list, m)
	  m = make(map[string]string)
	  m["room"] = "tanguy,scott,heidi,jenn"
	  m["id"] = "126"
	  list = append(list, m)
  }
  if t.Index == "1" {
	  m := make(map[string]string)
	  m["room"] = "payouts"
	  m["id"] = "127"
	  list = append(list, m)
  }

  return list
}

func GetTeams() []Team {
        teams := make([]Team, 0)

	slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
                team := Team{}
                team.Index = fmt.Sprintf("%d", i)
		team.Api = slack.New(os.Getenv(key))
		//fmt.Println(team)
                teams = append(teams, team)

		i++
		if i >= slack_teams {
			break
		}
	}
        return teams

}
