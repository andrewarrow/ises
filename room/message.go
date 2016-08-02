package room

import "os"

func WriteMessageToDisk(filename string, h map[string]string) {
	fstr := "cache/" + filename + "/" + h["time"] + "_" + h["who"]
	_, err := os.Stat(fstr)
	if os.IsNotExist(err) {
		f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(h["text"])
		f.Close()
	}
}
