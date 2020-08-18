package main

import (
	"fmt"
)

type Car struct {
	Model string
}

func main() {
	//make channel
	//<name> chan <data_type>
	c := make(chan int, 3)
	//send data in goroutings
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()
	//receive and print
	for i := range c {
		fmt.Println(i)
	}
	//string

	c1 := make(chan *Car, 3)
	go func() {
		c1 <- &Car{"1"}
		c1 <- &Car{"2"}
		c1 <- &Car{"3"}
		c1 <- &Car{"4"}
		close(c1)
	}()
	//receive and print
	for i := range c1 {
		fmt.Println(i.Model)
	}

}
