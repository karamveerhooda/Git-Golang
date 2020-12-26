/*Hands-on exercise #6
Build and use an anonymous func
code: https://play.golang.org/p/DQX3xEIcRe  */
package main

import "fmt"

func main() {

	func(a int) {
		fmt.Println("I am anoynymous func", a)
	}(33)

}
