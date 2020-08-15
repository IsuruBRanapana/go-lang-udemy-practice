package main

import (
	"fmt"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Pinger")
		time.Sleep(time.Second * 1)
		ponger <- 1
	}
}
func ponger(ponger <-chan int, pinger chan<- int) {
	for {
		<-ponger
		fmt.Println("Ponger")
		time.Sleep(time.Second * 1)
		pinger <- 1
	}
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go pinger(c1, c2)
	go ponger(c2, c1)

	c1 <- 1
	select {}
}
