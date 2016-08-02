package room

import "io/ioutil"
import "os"

func IdToString(id, team string) string {
	root := os.Getenv("ISES_ROOT")
	fstr := root + "/id_lookup/" + team + "/" + id
	b, err := ioutil.ReadFile(fstr)
	if err != nil {
		return "[NEW]"
	}
	return string(b)
}

func StringToId(name, team string) string {
	root := os.Getenv("ISES_ROOT")
	fstr := root + "/reverse_lookup/" + team + "/" + name
	b, err := ioutil.ReadFile(fstr)
	if err != nil {
		return "oops"
	}
	return string(b)
}
