//////////////////////////////Callback need to revisit after completion//////////////////////////////////////////////////////////////////

/*A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
code: https://play.golang.org/p/0yGYPKh1y7


*/
package main

import "fmt"

func main() {

	x := foo(bar())

	fmt.Println(x)
}

func bar() {
	fmt.Println("I am in bar")
}

func foo(f func()) {
	fmt.Println("I am in foo")
}
