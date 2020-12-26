// Unidirectional
package main

import (
	"fmt"
)

func main() {

	chanvar := make(<-chan int, 2)
	chanvar <- 10 // This will give error because
	chanvar <- 20 // Channel is receive only not send, and here
	//sending values the channel

	fmt.Println(<-chanvar)
	fmt.Println(<-chanvar)

	fmt.Printf("%T\t", chanvar)
}
