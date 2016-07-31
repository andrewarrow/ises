package main

import "fmt"
import "os"
import "github.com/nlopes/slack"
import "strconv"

func main() {
	slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
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
