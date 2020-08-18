package main

import (
	"fmt"
	"sync"
)

func main() {

	//wait groups
	//add,done,wait
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("finished")
	wg.Wait()
	fmt.Println("all done")

	//sync.mutex are use to lock and unlock resources
}
