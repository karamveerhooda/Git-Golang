//Channel can be single direction means only receivers or senders
//Channel can be bidrectional
package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan int, 2) // here chan int is bidirectional, means it can send and recieve
	//values from channel
	chanvar <- 10 // sending value to channel
	chanvar <- 20

	fmt.Println(<-chanvar) // Receiving values from channel
	fmt.Println(<-chanvar)

	fmt.Printf("%T\t", chanvar)
}
