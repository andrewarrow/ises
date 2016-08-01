package main

import "fmt"

//import "os"
//import "github.com/nlopes/slack"
//import "strconv"
import "flag"

var (
	ft     = flag.String("t", "0", "team index zero based")
	fid    = flag.String("id", "0", "id")
	fquick = flag.Bool("q", false, "quick mode")
	fsync  = flag.Bool("s", false, "sync mode")
	fread  = flag.Bool("r", false, "read mode")
)

func main() {
	flag.Parse()
	//slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	if *fquick == true {
		handleQuickMode()
		return
	}
	if *fsync == true {
		handleSyncMode()
		return
	}
	if *fread == true {
		handleReadMode()
		return
	}

	fmt.Println("end")
}
