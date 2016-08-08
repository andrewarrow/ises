package soeasy

import "time"
import "github.com/andrewarrow/ises/room"
import "strconv"
import "io/ioutil"

type RecentRoom struct {
	ts        int64
	fullName  string
	realId    string
	name      string
	team      string
	teamIndex int
}

func NewRecentRoom(full string) RecentRoom {
	rr := RecentRoom{}
	rr.fullName = full
	rr.ts = time.Now().Unix()
	rr.name = full[2:len(full)]
	rr.team = full[0:1]
	rr.teamIndex, _ = strconv.Atoi(rr.team)
	rr.realId = room.StringToId(rr.name, rr.team)
	return rr
}

func simpleMostRecent() []string {
	subfiles, _ := ioutil.ReadDir("cache/messages/")
	list := make([]string, 0)
	for _, sub := range subfiles {
		if sub.Name() == ".DS_Store" {
			continue
		}
		list = append(list, sub.Name())
	}
	if len(list) < 5 {
		panic("you need more rooms")
	}
	return list[0:5]
}

func mostRecent() []string {
	return computeLatestRooms()
}

func recentDefaults(sec *SoEasyClient) {

	for _, r := range mostRecent() {
		sec.addToRecentOrUpdateTs(NewRecentRoom(r))
	}
}
