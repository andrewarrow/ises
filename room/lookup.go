package room

import "io/ioutil"

func IdToString(id, team string) string {
	fstr := "id_lookup/" + team + "/" + id
	b, err := ioutil.ReadFile(fstr)
	if err != nil {
		return "[NEW]"
	}
	return string(b)
}

func StringToId(name, team string) string {
	fstr := "reverse_lookup/" + team + "/" + name
	b, err := ioutil.ReadFile(fstr)
	if err != nil {
		return "oops"
	}
	return string(b)
}
