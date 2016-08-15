package room

import "time"
import "fmt"

func (t Team) Sync(fullName, realId string) error {
	ts := time.Now().Unix() - int64(432000)
	missing := t.History(realId, realId[0:1], fmt.Sprintf("%d", ts))
	for _, h := range missing {
		WriteMessageToDisk(fullName, h)
	}
	return nil
}
