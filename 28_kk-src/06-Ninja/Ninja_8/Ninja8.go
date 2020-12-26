/*Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
code: https://play.golang.org/p/c2HwqVE1Rd
*/
package main

import (
	"fmt"
)

func main() {

	f := foo()

	fmt.Println(f())
}

func foo() func() int { // that means return a func that returns int

	return func() int { // these line tells return a function have return type int , so need to return type int with value 42
		return 42
	}
}
