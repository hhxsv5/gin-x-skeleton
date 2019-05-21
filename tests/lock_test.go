package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	go order()

	for {
		log.Println("xx")
		time.Sleep(time.Second)
	}
}

func order() {
	lock := &sync.RWMutex{}
	lock.Lock()
	log.Print("hello world.")
	lock.Lock()
	log.Print("hello 123.")
}
