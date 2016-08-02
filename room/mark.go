package room

import "io/ioutil"
import "fmt"
import "os"

func MarkRoomWithRedDot(name, team string) {
	alpha := []string{"a", "b", "c", "d", "e",
		"f", "g", "h"}
	// write out bash sh named a_name, b_name
	files, _ := ioutil.ReadDir("ui/")
	size := len(files)

	if size > 25 {
		return
	}

	fstr := fmt.Sprintf("ui/%s_%s", alpha[size], name)
	cmd := fmt.Sprintf("TEAM=%s ROOM=%s ../ises -r", team, name)
	f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0660)
	_, _ = f.WriteString(cmd)
	f.Close()

}
