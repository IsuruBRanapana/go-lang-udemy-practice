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

//pointers
func swap(m1,m2 *int){
	var temp int
	temp = *m2
	*m2 = *m1
	*m1=temp
}

func main(){
	todo()
	m1,m2 := 2,3
	fmt.Println(m1,m2)
	swap(&m1,&m2)
	fmt.Println(m1,m2)
}