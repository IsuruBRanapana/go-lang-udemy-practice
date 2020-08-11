package main

import "fmt"

func todo(){
	arr := [] int{2,3,4,5,6}
	fmt.Println(arr)
	arr2:=[] string{"my","name","is"}
	fmt.Println(arr2)
	arr2=append(arr2, "isuru","ranapana")
	fmt.Println(arr2)
}

func main(){
	todo()
}