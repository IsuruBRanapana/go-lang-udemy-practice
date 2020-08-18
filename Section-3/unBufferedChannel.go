package main

import (
	"fmt"
	"time"
)

func main() {
	//make channel
	//<name> chan <data_type>
	c := make(chan int)
	//send data in goroutings
	go func() {
		c <- 1
	}()
	//receive and print
	val := <-c
	fmt.Println(val)

	go func() {
		c <- 2
	}()
	time.Sleep(time.Second * 2)
	val = <-c
	println(val)
}
