package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Go-land")
	Anything(2.44)
	Anything("Isuru")
	Anything(struct{}{})

	mymap := make(map[string]interface{})
	mymap["name"] = "isuru"
	mymap["age"] = 25
	fmt.Println(mymap)
}
