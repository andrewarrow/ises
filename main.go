package main

import "github.com/andrewarrow/ises/soeasy"
import "os"

func log(str string) {
	fstr := "log.log"
	f, _ := os.OpenFile(fstr, os.O_APPEND|os.O_WRONLY, 0660)
	_, _ = f.WriteString(str + "\n")
	f.Close()
}

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
