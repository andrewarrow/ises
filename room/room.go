package room

import "github.com/nlopes/slack"
import "strconv"
import "os"
import "fmt"

type Team struct {
	Api   *slack.Client
	Rtm   *slack.RTM
	Index string
}

func (t Team) Say(id, say string) error {
	params := slack.PostMessageParameters{}
	params.AsUser = true
	//attachment := slack.Attachment{ Pretext: "", Text:    *say, }
	//params.Attachments = []slack.Attachment{attachment}
	_, _, err := t.Api.PostMessage(id, say, params)
	return err
}

func (t Team) History(id, thing string) []map[string]string {
	list := make([]map[string]string, 0)
	hp := slack.HistoryParameters{Oldest: "", Latest: "", Count: 10, Inclusive: false, Unreads: false}
	if thing == "c" {
		history, _ := t.Api.GetChannelHistory(id, hp)
		//fmt.Println("ch ", err)
		for _, message := range history.Messages {
			m := make(map[string]string)
			m["text"] = message.Msg.Text
			m["time"] = message.Msg.Timestamp
			m["who"] = message.Msg.User
			list = append(list, m)
		}
	} else if thing == "g" {
		history, _ := t.Api.GetGroupHistory(id, hp)
		//fmt.Println("gh ", err)
		for _, message := range history.Messages {
			m := make(map[string]string)
			m["text"] = message.Msg.Text
			m["time"] = message.Msg.Timestamp
			m["who"] = message.Msg.User
			list = append(list, m)
		}
	} else if thing == "i" {
		history, _ := t.Api.GetIMHistory(id, hp)
		//fmt.Println("api.GetIMHistory ", err)
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

	ims, err := t.Api.GetIMChannels()
	fmt.Println("getting ims ", err)
	for _, im := range ims {
		m := make(map[string]string)
		m["room"] = IdToString(im.User, t.Index)
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
		team.Rtm = team.Api.NewRTM()
		//fmt.Println(team)
		teams = append(teams, team)

		i++
		if i >= slack_teams {
			break
		}
	}
	return teams

}
