package room

import "sync"
import "os"

var WriteFileMutex = &sync.Mutex{}

func WriteMessageToDisk(filename string, h map[string]string) {
	dir := "cache/messages/" + filename + "/"
	_ = os.MkdirAll(dir, os.ModePerm)
	fstr := dir + h["time"] + "_" + h["who"]
	_, err := os.Stat(fstr)
	if os.IsNotExist(err) {
		WriteFileMutex.Lock()
		f, _ := os.OpenFile(fstr, os.O_CREATE|os.O_WRONLY, 0660)
		_, _ = f.WriteString(h["text"])
		f.Close()
		WriteFileMutex.Unlock()
	}
}
