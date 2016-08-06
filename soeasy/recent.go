package soeasy

import "time"

type RecentRoom struct {
	ts       int64
	fullName string
	realId   string
	name     string
	team     string
}

func NewRecentRoom(full string) RecentRoom {
	rr := RecentRoom{}
	rr.fullName = full
	rr.ts = time.Now().Unix()
	rr.name = full[2:len(full)]
	rr.team = full[0:1]
	return rr
}

func recentDefaults() []RecentRoom {
	rd := make([]RecentRoom, 0)

	rd = append(rd, NewRecentRoom("1_for_andrew"))
	rd = append(rd, NewRecentRoom("0_aa"))
	rd = append(rd, NewRecentRoom("1_general"))
	rd = append(rd, NewRecentRoom("2_jasoncarulli"))
	rd = append(rd, NewRecentRoom("3_ast"))
	rd = append(rd, NewRecentRoom("3_office"))
	return rd
}
