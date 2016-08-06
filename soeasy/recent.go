package soeasy

import "time"
import "github.com/andrewarrow/ises/room"
import "strconv"

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

func recentDefaults(sec *SoEasyClient) {

	sec.addToRecentOrUpdateTs(NewRecentRoom("1_for_andrew"))
	sec.addToRecentOrUpdateTs(NewRecentRoom("0_aa"))
	sec.addToRecentOrUpdateTs(NewRecentRoom("1_general"))
	sec.addToRecentOrUpdateTs(NewRecentRoom("2_jasoncarulli"))
	sec.addToRecentOrUpdateTs(NewRecentRoom("3_ast"))
	sec.addToRecentOrUpdateTs(NewRecentRoom("3_office"))
}
