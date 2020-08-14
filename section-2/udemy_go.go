package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings
	m1 := "My name"
	m2 := "name"
	fmt.Println(m1)
	//check sub Strings
	fmt.Println(strings.Contains(m1,m2))
	fmt.Println(strings.ReplaceAll(m2,"ame","oun"))
	fmt.Println(strings.Split(m1," "))
}