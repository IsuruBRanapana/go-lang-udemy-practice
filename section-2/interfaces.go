package main

import "fmt"

type Car1 interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Navara struct {
	NavaraModel string
}

//if the lambo not satisfied all methods in interface this one is not compiling
func NewModel(arg string) Car1 {
	return &Lambo{arg}
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo is on move")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop() {
	fmt.Println("Lambo is stopping")
}

func (n *Navara) Drive() {
	fmt.Println("Navara is on move")
	fmt.Println(n.NavaraModel)
}
func main() {
	l := NewModel("Gallardo")
	n := Navara{"Navara"}
	l.Drive()
	n.Drive()
}
