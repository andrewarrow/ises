package main

import "fmt"
import "os"
import "github.com/nlopes/slack"
import "strconv"
import "flag"

var (
	team   = flag.String("t", "0", "team index zero based")
	id     = flag.String("id", "0", "id")
	people = flag.String("p", "0", "list people")
	say    = flag.String("s", "0", "say")
)

func main() {
	flag.Parse()
	slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)

	if *say != "0" {
		key := fmt.Sprintf("SLACK_TOKEN_%s", *team)
		api := slack.New(os.Getenv(key))
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Pretext: "2some pretext",
			Text:    "3some text",
		}
		params.Attachments = []slack.Attachment{attachment}
		_, _, _ = api.PostMessage(*id, "1Some text", params)

		return
	}

	if *people != "0" {
		i := int64(0)
		for {
			key := fmt.Sprintf("SLACK_TOKEN_%d", i)
			api := slack.New(os.Getenv(key))
			users, _ := api.GetUsers()
			for _, user := range users {
				fmt.Println(user.ID, i, user.Name)
			}

			i++
			if i >= slack_teams {
				break
			}
		}
		return
	}

	if *id != "0" {
		key := fmt.Sprintf("SLACK_TOKEN_%d", 0)
		api := slack.New(os.Getenv(key))
		hp := slack.HistoryParameters{Oldest: "", Latest: "", Count: 10, Inclusive: false, Unreads: false}
		h, err := api.GetChannelHistory(*id, hp)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println(h)
		return
	}

	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
		fmt.Println(key)
		api := slack.New(os.Getenv(key))
		channels, err := api.GetChannels(false)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		for _, channel := range channels {
			fmt.Println(channel.ID, channel.Name)
		}

		i++
		if i >= slack_teams {
			break
		}
	}

}
