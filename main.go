package main

import "fmt"
import "flag"
import "os"

var (
	IsesRoot string
	fsync    = flag.Bool("s", false, "sync mode")
	fread    = flag.Bool("r", false, "read mode")
	fdaemon  = flag.Bool("d", false, "daemon mode")
	flookup  = flag.Bool("l", false, "fill lookup data")
)

func main() {
	flag.Parse()
	if *fsync == false && *fread == false && *fdaemon == false && *flookup == false {
		fmt.Println("use --help")
		return
	}

	IsesRoot = os.Getenv("ISES_ROOT")

	if *fsync == true {
		handleSyncMode()
		return
	}
	if *fread == true {
		handleReadMode()
		return
	}

	if *fdaemon == true {
		handleDaemonMode()
		return
	}

	if *flookup == true {
		handleLookupMode()
	}

	fmt.Println("end")
}
