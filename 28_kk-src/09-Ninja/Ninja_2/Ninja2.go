package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) speak() string {
	p.first = "Hello I am multilingual"
	return p.first
}
func (p person) native() string {
	p.first = "Indian"
	return p.first
}

type human interface {
	speak() string
	native() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
	fmt.Println(h.native())
}

func main() {

	p1 := person{
		first: "karam",
	}

	saySomething(&p1)

	fmt.Println("hello baby")

}
