package ch01

import (
	"fmt"
	"log"
	"os"
	"path"
)

func ch01CustomLogLineNumber() {
	logfile := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(logfile)
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	lstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum", lstdFlags)
	iLog.Println("Mastering Go")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log")
}
