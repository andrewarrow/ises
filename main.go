package main

import "github.com/andrewarrow/ises/soeasy"
import "os"
import "fmt"

func main() {
	args := os.Args[1:]
	initMode := false
	_, err := os.Stat("cache")
	if os.IsNotExist(err) {
		initMode = true
	}

	if len(args) == 1 && args[0] == "--init" {
		initMode = true
	}

	if initMode {
		fmt.Println("Starting init mode...")
		soeasy.SoEasySetup()
		fmt.Println("\nSetup complete, run ./ises again")
	} else {
		var client *soeasy.SoEasyClient
		client = soeasy.NewSoEasyClient()
		client.Paint()
		client.InputLoop()
	}
}
