package main

import "fmt"

//import "os"
//import "github.com/nlopes/slack"
//import "strconv"
import "flag"

var (
	ft  = flag.String("t", "0", "team index zero based")
	fid = flag.String("id", "0", "id")
)

func main() {
	flag.Parse()
	//slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	fmt.Println("")
}
