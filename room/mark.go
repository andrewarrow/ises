package room

import "io/ioutil"
import "fmt"
import "os"

func MarkRoomWithRedDot(name, team string) {
	alpha := []string{"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// write out bash sh named a_name, b_name
	files, _ := ioutil.ReadDir("ui/")
	size := len(files)

	if size > 25 {
		return
	}

	for _, a := range alpha {
		file := fmt.Sprintf("ui/%s_%s_%s", a, team, name)
		_, err := os.Stat(file)
		if !os.IsNotExist(err) {
			return
		}
	}

	fstr := fmt.Sprintf("ui/%s_%s_%s", alpha[size], team, name)
	cmd := fmt.Sprintf("TEAM=%s ROOM=%s ../ises -r", team, name)
	f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0755)
	_, _ = f.WriteString(cmd)
	f.Close()

}
