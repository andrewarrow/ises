package room

type RoomFinder struct {
}

func Recents() []int {
	s := NewSession()
	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
		fmt.Println(key)
		api := slack.New(os.Getenv(key))

		i++
		if i >= slack_teams {
			break
		}
	}

}
