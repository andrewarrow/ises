package main

import "fmt"
import "flag"

var (
	fsync   = flag.Bool("s", false, "sync mode")
	fread   = flag.Bool("r", false, "read mode")
	fdaemon = flag.Bool("d", false, "daemon mode")
	flookup = flag.Bool("l", false, "fill lookup data")
)

func main() {
	flag.Parse()
	if *fsync == false && *fread == false && *fdaemon == false && *flookup == false {
		fmt.Println("use --help")
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

	if *fdaemon == true {
		handleDaemonMode()
		return
	}

	if *flookup == true {
		handleLookupMode()
	}

	fmt.Println("end")
}
