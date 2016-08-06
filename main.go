package main

import "github.com/andrewarrow/ises/soeasy"

/*
type RecentRoom struct {
	ts   int64
	name string
}

type ByRoomAge []RecentRoom

func (a ByRoomAge) Len() int           { return len(a) }
func (a ByRoomAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRoomAge) Less(i, j int) bool { return a[i].ts < a[j].ts }


func findRecents() []RecentRoom {
	list := make([]RecentRoom, 0)
	for k, v := range recent {
		rr := RecentRoom{}
		rr.ts = v
		rr.name = k
		list = append(list, rr)
	}
	sort.Sort(ByRoomAge(list))
	return list

}

func log(str string) {
	fstr := "log.log"
	f, _ := os.OpenFile(fstr, os.O_APPEND|os.O_WRONLY, 0660)
	_, _ = f.WriteString(str + "\n")
	f.Close()
}
*/

func main() {
	/*
		args := os.Args[1:]
		if len(args) == 0 {
			fmt.Println("./ises rid")
			return
		}
	*/
	var client *soeasy.SoEasyClient
	client = soeasy.NewSoEasyClient()
	client.Paint()
	client.InputLoop()
}
