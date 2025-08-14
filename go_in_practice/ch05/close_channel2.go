package ch05

import (
	"log"
	"time"
)

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			log.Println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	msg := make(chan string)
	done := make(chan bool)
	go send(msg, done)
	for {
		select {
		case m := <-msg:
			log.Println(m)
		case <-time.After(5 * time.Second):
			done <- true
			return
		}
		
	}
}
