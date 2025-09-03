package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

var directory string
var watcher *fsnotify.Watcher

type FileUpdateInfo struct {
	Name      string `json:"name"`
	Op        string `json:"operation"`
	SizeBytes int    `json:"size"`
}

func init() {
	if osdir := os.Getenv("SSE_DIRECTORY"); osdir == "" {
		panic("SSE_DIRECTORY environment variable not set")
	} else {
		directory = osdir
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request)

func main() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	watcher.Add(directory)

	http.HandleFunc("/sse", sseHandler)
	http.HandleFunc("/files", fileesHandler)
	http.ListenAndServe(":8080", nil)
}
