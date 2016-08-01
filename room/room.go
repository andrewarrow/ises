package room

import "github.com/nlopes/slack"
import "strconv"
import "os"
import "fmt"

type Team struct {
	Api   *slack.Client
	Index string
}

func (t Team) Recents() []map[string]string {
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

func (t Team) History(id, thing, latest string) []map[string]string {
	list := make([]map[string]string, 0)
	hp := slack.HistoryParameters{Oldest: "", Latest: latest, Count: 10, Inclusive: false, Unreads: false}
	if thing == "c" {
		history, err := t.Api.GetChannelHistory(id, hp)
		fmt.Println("ch ", err)
		for _, message := range history.Messages {
			m := make(map[string]string)
			m["text"] = message.Msg.Text
			m["time"] = message.Msg.Timestamp
			m["who"] = message.Msg.User
			list = append(list, m)
		}
	} else if thing == "g" {
		history, err := t.Api.GetGroupHistory(id, hp)
		fmt.Println("gh ", err)
		for _, message := range history.Messages {
			m := make(map[string]string)
			m["text"] = message.Msg.Text
			m["time"] = message.Msg.Timestamp
			m["who"] = message.Msg.User
			list = append(list, m)
		}
	} else if thing == "i" {
		history, err := t.Api.GetIMHistory(id, hp)
		fmt.Println("api.GetIMHistory ", err)
		for _, message := range history.Messages {
			m := make(map[string]string)
			m["text"] = message.Msg.Text
			m["time"] = message.Msg.Timestamp
			m["who"] = message.Msg.User
			list = append(list, m)
		}
	}
	return list
}

func (t Team) Rooms() []map[string]string {
	list := make([]map[string]string, 0)

	channels, err := t.Api.GetChannels(false)
	fmt.Println("getting channels ", err)
	for _, channel := range channels {
		m := make(map[string]string)
		m["room"] = channel.Name
		m["id"] = channel.ID
		m["thing"] = "c"
		list = append(list, m)
	}
	groups, err := t.Api.GetGroups(false)
	fmt.Println("getting groups ", err)
	for _, group := range groups {
		m := make(map[string]string)
		m["room"] = group.Name
		m["id"] = group.ID
		m["thing"] = "g"
		list = append(list, m)
	}

	userMap := make(map[string]string)
	users, err := t.Api.GetUsers()
	fmt.Println("getting users ", err)
	for _, user := range users {
		userMap[user.ID] = user.Name
	}
	ims, err := t.Api.GetIMChannels()
	fmt.Println("getting ims ", err)
	for _, im := range ims {
		m := make(map[string]string)
		m["room"] = userMap[im.User]
		m["id"] = im.ID
		m["thing"] = "i"
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