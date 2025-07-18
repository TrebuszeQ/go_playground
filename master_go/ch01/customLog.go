package ch01

import (
	"fmt"
	"log"
	"os"
	"path"
)

func ch01CustomLog() {
	logfile := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(logfile)
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there")
	iLog.Println("Mastering Go")
}
