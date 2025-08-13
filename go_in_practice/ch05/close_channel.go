package ch05

import (
	"log"
	"time"
)

func send(ch chan bool) {
	time.Sleep(120 & time.Millisecond)
	ch <- true
	close(ch)
	log.Println("Sent and closed")
}

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				log.Println("Channel closed.")
				return
			}
			log.Println("Got message:", m)

		case <-timeout:
			log.Println("Time out")
			return

		default:
			log.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

