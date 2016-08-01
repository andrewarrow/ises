package room

import "io/ioutil"

func IdToString(id, team string) string {
	fstr := "id_lookup/" + team + "/" + id
	b, _ := ioutil.ReadFile(fstr)
	return string(b)
}
