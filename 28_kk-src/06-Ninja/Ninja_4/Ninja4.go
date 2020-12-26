/*Hands-on exercise #4
Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age
	attach a method to type person with
		the identifier “speak”  -- it means func (p person) speak() line 42
		the method should have the person say their name and age
create a value of type person  -- here values are p1 nad p2
call the method from the value of type person
code: https://play.golang.org/p/NnXyWdqbbw*/
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		"karam",
		"hooda",
		39,
	}
	p2 := person{
		"punam",
		"khokhar",
		39,
	}
	fmt.Println(p1, p2)
	p1.speak()
	p2.speak()

}

func (p person) speak() {

	fmt.Println(p.age, p.first, p.last)
}
